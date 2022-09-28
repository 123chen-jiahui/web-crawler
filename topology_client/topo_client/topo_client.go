package topo_client

import (
	"context"
	"fmt"

	pb "github.com/proto"

	"google.golang.org/grpc"
)

func format(mf map[string][]string) (int32, []string, map[int32]*pb.Destination) {
	m_check := make(map[string]int32)    // 检查所有url是否重复
	contents := make([]string, 0)        // 返回url数组
	g := make(map[int32]*pb.Destination) // 返回的图
	var count int32                      // 统计一共有几个结点

	for k, v := range mf {
		var index_from int32
		if i, ok := m_check[k]; ok {
			index_from = int32(i)
		} else {
			index_from = count
			contents = append(contents, k)
			m_check[k] = index_from
			count += 1
		}

		arr := make([]int32, 0)
		for _, url := range v {
			var index_to int32
			if i, ok := m_check[url]; ok {
				index_to = int32(i)
			} else {
				index_to = count
				contents = append(contents, url)
				m_check[url] = index_to
				count += 1
			}
			arr = append(arr, index_to)
		}

		g[index_from] = &pb.Destination{
			Nodes: arr,
		}
	}
	return count, contents, g
}

func SendTopoRequest(mf map[string][]string) ([]string, string) {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewDrawClient(conn)

	total, contents, g := format(mf)

	r, err := c.Topology(context.Background(), &pb.TopologyRequest{
		SaveLocation: "./",
		Total:        total,
		Contents:     contents,
		G:            g,
	})
	if err != nil {
		panic(err)
	}
	if r.Ok {
		fmt.Println("Ok")
	}

	return contents, r.FilePath
}
