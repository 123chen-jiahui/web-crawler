package fetch

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	NMatch   = 2
	NGoal    = 3
	MaxDepth = 3
)

// var Pattern = []string{"tongji", "edu", "cn"}

// type Url_Depth struct {
// 	url   string
// 	depth int
// }

type Fetcher interface {
	Fetch(url string) (err error)
}

type MyFetcher map[string][]string

func (mf *MyFetcher) Fetch(url string, pattern []string) ([]string, error) {
	body := GetHtml(url)
	urls, ok := GetURL(body, pattern)
	if ok {
		return urls, nil
	} else {
		return urls, errors.New("crawler failed: no urls in html")
	}
}

type FetchState struct {
	Mu      sync.Mutex
	Fetched map[string]bool // 并发不安全，需要锁来保护
}

func Crawl(url string) (int, int, string, map[string][]string) {
	// 解析pattern
	pattern := strings.Split(url, ".")
	pattern = pattern[1:] // 将http(s)://www去掉

	mf := make(MyFetcher, 0)
	f := FetchState{
		Fetched: make(map[string]bool),
	}
	ConcurrencyMutex(url, pattern, 0, mf, &f)

	// 找到入度最多的url并返回
	var nedge int
	var nnode int
	m := make(map[string]int) // 统计
	for _, v := range mf {
		nedge += len(v)
		for _, url := range v {
			if m[url] == 0 {
				nnode += 1
			}
			m[url] += 1
		}
	}
	for k := range mf {
		if m[k] == 0 {
			nnode += 1
		}
	}
	var res string
	var max int
	for k, v := range m {
		if v > max {
			max = v
			res = k
		}
	}
	return nnode, nedge, res, map[string][]string(mf)
}

func ConcurrencyMutex(url string, pattern []string, depth int, mf MyFetcher, f *FetchState) {
	if depth >= MaxDepth {
		return
	}

	f.Mu.Lock()
	already := f.Fetched[url]
	f.Fetched[url] = true
	f.Mu.Unlock()

	if already {
		return
	}

	urls, err := mf.Fetch(url, pattern)
	if err != nil {
		return
	}
	mf[url] = urls
	// 可以将done理解为一个条件变量，具体内容可以复习https://github.com/123chen-jiahui/6.S081notes/blob/main/sleep%26wakeup.md
	var done sync.WaitGroup
	for _, u := range urls {
		done.Add(1)
		// 注意此处，不能写成无参形式，否则u会变化
		go func(u string) {
			defer done.Done()
			ConcurrencyMutex(u, pattern, depth+1, mf, f)
		}(u)
	}
	done.Wait()
}

// func worker(u_d Url_Depth, ch chan []Url_Depth, mf MyFetcher) {
// 	urls, err := mf.Fetch(u_d.url)

// 	if err != nil {
// 		ch <- []Url_Depth{}
// 	} else {
// 		res := make([]Url_Depth, len(urls))
// 		for i, _ := range res {
// 			res[i].url = urls[i]
// 			res[i].depth = u_d.depth + 1
// 		}
// 		ch <- res
// 	}
// }

// func Coordinator(ch chan []Url_Depth, mf MyFetcher) {
// 	n := 1 // n表示正在进行的task数量
// 	fetched := make(map[string]bool)
// 	for us_d := range ch { // 如果channel中没有东西，会一直等待
// 		for _, u_d := range us_d {
// 			if fetched[u_d.url] == false && u_d.depth <= MaxDepth {
// 				fetched[u_d.url] = true
// 				n += 1

// 				go worker(Url_Depth{url: u_d.url, depth: u_d.depth + 1}, ch, mf)
// 			}
// 		}
// 		n -= 1
// 		if n == 0 {
// 			break
// 		}
// 	}
// }

func GetHtml(url string) string {
	res, err := http.Get(url)
	if err != nil {
		// fmt.Println(err)
		return "something wrong"
	}
	body, _ := ioutil.ReadAll(res.Body) //转换byte数组
	defer res.Body.Close()
	//io.Copy(os.Stdout, res.Body)//写到输出流，
	bodystr := string(body) //转换字符串
	return bodystr
}

func check_pattern(url string, pattern []string) bool {
	var count int
	for _, u := range pattern {
		if strings.Contains(url, u) {
			count += 1
		}
	}
	if count > NMatch {
		return false
	} else {
		return true
	}
}

func GetURL(body string, pattern []string) ([]string, bool) {
	regex := `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`
	compileRegex := regexp.MustCompile(regex)
	Arr := compileRegex.FindAllString(body, -1)
	m_url := make(map[string]bool) // 判断url是否重复
	m_index := make(map[int]bool)  // 判断index是否重复
	var matchArr []string          // 保存url
	var count int                  // 计数器
	// 消除不符合条件的
	// 随机选取
	rand.Seed(time.Now().Unix())
	for {
		// 所有url都找了一遍，没有达到指定个数
		if len(m_index) == len(Arr) {
			break
		}

		index := rand.Intn(len(Arr))
		_, ok := m_index[index]
		if !ok {
			m_index[index] = true
		} else {
			continue
		}

		u := Arr[index]
		if check_pattern(u, pattern) {
			_, ok := m_url[u]
			if !ok {
				matchArr = append(matchArr, u)
				m_url[u] = true
				count += 1
				if count >= NGoal {
					break
				}
			}
		}
	}
	return matchArr, len(matchArr) > 0
}
