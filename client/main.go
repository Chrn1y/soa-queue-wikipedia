package main

import (
	"context"
	"fmt"
	wikipedia_proto "github.com/Chrn1y/soa-queue-wikipedia/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	addr = "127.0.0.1"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial(addr+":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	c := wikipedia_proto.NewAppClient(conn)
	for {
		input := ""
		fmt.Scanln(&input)
		switch input {
		case "0":
			return
		case "1":
			from, to := "", ""
			fmt.Scanln(&from)
			fmt.Scanln(&to)
			process, err := c.Process(ctx, &wikipedia_proto.Request{
				Link1: from,
				Link2: to,
			})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(process.Id)
		case "2":
			id := ""
			fmt.Scanln(&id)
			get, err := c.Get(ctx, &wikipedia_proto.Id{Id: id})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(get.String())
		}
	}

}
