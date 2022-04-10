package main

import (
	"encoding/json"
	"fmt"
	"github.com/Chrn1y/soa-queue-wikipedia/wikipedia"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strconv"
)

const (
	rabbitmq = "amqp://guest:guest@localhost:5672/"
)

func main() {
	raw := os.Getenv("WORKER_COUNT")
	if raw == "" {
		log.Fatal("no WORKER_COUNT was provided")
	}
	num, err := strconv.Atoi(raw)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := amqp.Dial(rabbitmq)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"worker-input", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	for range make([]struct{}, num) {
		go func() {
			for d := range msgs {
				log.Println("Received a message")
				inp := &wikipedia.Inp{}
				err = json.Unmarshal(d.Body, inp)
				if err != nil {
					log.Println(err)
					continue
				}
				process, _ := wikipedia.Process(inp)
				out, err := json.MarshalIndent(process, "", "\t")
				if err != nil {
					log.Println(err)
					continue
				}
				err = os.WriteFile(fmt.Sprintf("../results/%s", process.Id), out, 0600)
				if err != nil {
					log.Print(err)
				}
			}
		}()
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
