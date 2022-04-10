package worker

import (
	"encoding/json"
	"github.com/Chrn1y/soa-queue-wikipedia/wikipedia"
	"github.com/streadway/amqp"
	"log"
)

type Results struct {
	Results map[string]*wikipedia.Result
}

type Closer func()

func Start(rabbitmq, queueName string, num int) (*Results, Closer) {
	conn, err := amqp.Dial(rabbitmq)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(
		queueName, // name
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
	results := &Results{Results: make(map[string]*wikipedia.Result)}
	for range make([]struct{}, num) {
		go func() {
			log.Println("worker started")
			for d := range msgs {
				log.Println("Received a message")
				inp := &wikipedia.Inp{}
				err = json.Unmarshal(d.Body, inp)
				if err != nil {
					log.Println(err)
					continue
				}
				process, _ := wikipedia.Process(inp)
				results.Results[process.Id] = process
				//out, err := json.MarshalIndent(process, "", "\t")
				//if err != nil {
				//	log.Println(err)
				//	continue
				//}
				//err = os.WriteFile(fmt.Sprintf("../results/%s", process.Id), out, 0600)
				//if err != nil {
				//	log.Print(err)
				//}
			}
		}()
	}
	log.Println("waiting for workers to start...")
	//for i := range start {
	//	<-start[i]
	//}
	return results, func() {
		conn.Close()
		ch.Close()
	}
}