package wikipedia

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

const pref = "/wiki/"

type Inp struct {
	Id   string
	From string
	To   string
}

type Result struct {
	Id   string
	Path []string
	Num  int32
}

func Process(inp *Inp) (*Result, error) {
	//path, ok := dfs(inp.From, inp.To, 0)
	//if !ok {
	//	return &Result{
	//		Id:   inp.Id,
	//	}, fmt.Errorf("no path was found for %s", inp.Id)
	//}
	path := bfs(inp.From, inp.To)
	return &Result{
		Id:   inp.Id,
		Path: path,
		Num: int32(len(path)),
	}, nil
}

const maxPath = 3
//var processed = make(map[string]struct{})
func dfs(inp string, to string, cur int32) ([]string, bool) {
	if cur >= maxPath{
		return nil, false
	}
	links := getLinks(inp)
	if strings.Contains(inp, "Trade_name"){
		log.Println(links)
	}
	for _, link := range links {
		if strings.Contains(to, link) || strings.Contains(link, to){
			return []string{inp, link}, true
		}
	}

	for _, link := range links {
		path, ok  := dfs(link, to, cur + 1)
		if !ok {
			continue
		}
		path = append(path, inp)
		return path, true
		}

	return nil, false
}

func getLinks(inp string) []string {
	resp, err := http.Get(inp)
	if err != nil {
		log.Fatal(err)
	}
	var links []string
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" && len(attr.Val) >= 6 && attr.Val[:6] == pref {
						links = append(links, "https://en.wikipedia.org" + attr.Val)
					}
				}
			}
		}
	}
}


type bfsNode struct {
	link string
	path []string
	next []*bfsNode
}

func bfs(inp string, to string) []string {
	q := []*bfsNode{
		{
			link: inp,
			path: []string{inp},
		},
	}
	//visited := []*Node{}
	for len(q) > 0 {
		vertex := q[0]
		//node, level := vertex.node, vertex.level
		//visited = append(visited, node)
		q = q[1:] //dequeue first node in queue(fifo)
		for _, link := range getLinks(vertex.link){
			if strings.Contains(link, to){
				return append(vertex.path, link)
			}
			q = append(q, &bfsNode{
				link: link,
				path:  append(vertex.path, link),
				next: nil,
			})
		}
		//if node.Left != nil{ //have both left and right since it's a perfect binary tree
		//	leftNode := NodeWithLevel{
		//		node: node.Left,
		//		level: level+1,
		//	}
		//	q = append(q, leftNode) //append left-child to back of queue(fifo)
		//
		//	rightNode := NodeWithLevel{
		//		node: node.Right,
		//		level: level+1,
		//	}
		//	q = append(q, rightNode) //append right-child to back of queue(fifo)
		//}
	}
	return nil
}