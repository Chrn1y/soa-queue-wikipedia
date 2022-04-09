package main

import (
	"github.com/Chrn1y/soa-queue-wikipedia/wikipedia"
	"log"
)

//var inp = "https://en.wikipedia.org/wiki/Popular_Astronomy_(US_magazine)"

func main() {

	process, err := wikipedia.Process(&wikipedia.Inp{
		Id:   "123123123",
		From: "https://en.wikipedia.org/wiki/Sonic_Team",
		//To:   "https://en.wikipedia.org/wiki/Pachinko",
		//To: "https://en.wikipedia.org/wiki/Pseudonym",
		To: "https://en.wikipedia.org/wiki/Communist_party",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(process.Path)
}
