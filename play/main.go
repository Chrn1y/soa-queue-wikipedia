package main

import (
	"log"
	"net/http"
)

var inp = "https://en.wikipedia.org/wiki/Popular_Astronomy_(US_magazine)"

func main() {

	resp, err := http.Get(inp)
	if err != nil {
		log.Fatal(err)
	}
	for k := range resp.Header {
		for i := range resp.Header[k]{
			println(k, resp.Header[k][i])
		}
	}
	println(resp.Request.URL.String())
	//println(resp.Header)
	//io.ReadAll(resp.Body)
	//inp := &wikipedia.Inp{
	//	From: "https://en.wikipedia.org/wiki/Galaxy",
	//	To: "https://en.wikipedia.org/wiki/Maximinus_Thrax",
	//}
	//out, err := wikipedia.Process(inp)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(out.Path)
}
