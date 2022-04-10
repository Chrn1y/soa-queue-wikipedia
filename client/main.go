package main

import (
	"context"
	"fmt"
	wikipedia_proto "github.com/Chrn1y/soa-queue-wikipedia/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	fmt.Println("Введите адрес сервера:")
	addr := ""
	fmt.Scanln(&addr)
	ctx := context.Background()
	conn, err := grpc.Dial(addr+":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := wikipedia_proto.NewAppClient(conn)
	for {
		fmt.Println("Для отправки запроса нажмите 1\nДля проверки наличия результата нажмите 2\nДля выхода нажмите 0")
		input := ""
		fmt.Scanln(&input)
		switch input {
		case "0":
			return
		case "1":
			fmt.Println("Введите ссылки:")
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
			fmt.Println("ID запроса", process.Id)
		case "2":
			fmt.Println("Введите ID запроса")
			id := ""
			fmt.Scanln(&id)
			get, err := c.Get(ctx, &wikipedia_proto.Id{Id: id})
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println("Длина пути:", get.Len, "\nПуть:")
				for _, link := range get.Links{
					fmt.Println(link)
				}
			}
		}
	}

}
