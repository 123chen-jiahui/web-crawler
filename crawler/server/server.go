package main

import (
	"fmt"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/fetch"
	"github.com/topo_client"
)

// ============================

type FetchWrap struct {
	NNode int
	NEdge int
	Most  string
	Seed  string
	Urls  []string
	Path  string
	Cost  time.Duration
}

type DemoWrap struct {
	Content []FetchWrap
	mu      sync.Mutex
}

// ============================

func main() {
	fs := http.FileServer(http.Dir("wwwroot"))
	http.Handle("/", fs)
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("wwwroot/index.html")
		if t == nil {
			panic("我草")
		}
		t.Execute(w, nil)
		r.ParseForm()
		// fmt.Fprintln(w, r.Form)
	})

	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("wwwroot/demo.html")
		if t == nil {
			fmt.Println(err)
			panic("我草")
		}

		var wg sync.WaitGroup
		var urls = []string{
			"https://www.tongji.edu.cn",
			"https://www.fudan.edu.cn",
			"https://www.sjtu.edu.cn",
			"https://www.mit.edu",
		}

		demo := DemoWrap{
			Content: []FetchWrap{},
		}

		for i, url := range urls {
			wg.Add(1)
			go func(url string, i int) {
				defer wg.Done()
				start := time.Now()
				nnode, nedge, most, mf := fetch.Crawl(url)

				// ==========================
				urls, path := topo_client.SendTopoRequest(mf)
				demo.mu.Lock()
				demo.Content = append(demo.Content, FetchWrap{
					NNode: nnode,
					NEdge: nedge,
					Most:  most,
					Seed:  url,
					Urls:  urls,
					Path:  path,
					Cost:  time.Since(start),
				})
				fmt.Println(demo.Content)
				demo.mu.Unlock()
				// ==========================
			}(url, i)
		}
		wg.Wait()
		t.Execute(w, demo.Content)
	})

	http.HandleFunc("/crawler", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		r.ParseForm()
		t, _ := template.ParseFiles("wwwroot/res.html")
		if t == nil {
			panic("我草")
		}
		url := r.Form["url"][0]

		nnode, nedge, most, mf := fetch.Crawl(url)

		// ==========================
		wrap := FetchWrap{
			NNode: nnode,
			NEdge: nedge,
			Most:  most,
			Seed:  url,
			Cost:  time.Since(start),
		}
		wrap.Urls, wrap.Path = topo_client.SendTopoRequest(mf)
		// ==========================

		t.Execute(w, wrap)
	})
	http.ListenAndServe(":8080", nil)
}
