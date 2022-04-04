package wikipedia

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
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
	return &Result{
		Id:   inp.Id,
		Path: getLinks(inp.From),
		Num:  1,
	}, nil
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
					if attr.Key == "href" && attr.Val[:6] == pref {
						links = append(links, "https://en.wikipedia.org" + attr.Val)
					}
				}
			}
		}
	}
}
