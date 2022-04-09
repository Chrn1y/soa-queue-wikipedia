package main

import (
	"context"
	"encoding/json"
	"fmt"
	wikipedia_proto "github.com/Chrn1y/soa-queue-wikipedia/proto"
	"github.com/Chrn1y/soa-queue-wikipedia/wikipedia"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"hash/fnv"
	"log"
	"net"
	"os"
)

const (
	rabbitmq  = "amqp://guest:guest@localhost:5672/"
	queueName = "worker-input"
)

type Impl struct {
	wikipedia_proto.UnimplementedAppServer
	ch *amqp.Channel
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprintf("%d", h.Sum32())
}

func (i *Impl) Process(ctx context.Context, req *wikipedia_proto.Request) (*wikipedia_proto.Id, error) {

	id := hash(req.Link1 + req.Link2)

	inp := &wikipedia.Inp{
		Id:   id,
		From: req.Link1,
		To:   req.Link2,
	}

	inpBytes, err := json.Marshal(inp)
	if err != nil {
		return nil, err
	}
	err = i.ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        inpBytes,
		})
	if err != nil {
		return nil, err
	}
	return &wikipedia_proto.Id{Id: id}, nil
}

func (i *Impl) Get(ctx context.Context, id *wikipedia_proto.Id) (*wikipedia_proto.Response, error) {
	res, err := os.ReadFile(fmt.Sprintf("../results/%s", id.Id))
	if err != nil {
		return nil, err
	}
	temp := &wikipedia.Result{}
	err = json.Unmarshal(res, temp)
	if err != nil {
		return nil, err
	}

	return &wikipedia_proto.Response{
		Links: temp.Path,
		Len:   temp.Num,
	}, nil
}

func main() {
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

	_, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	log.Println("queue declared")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	wikipedia_proto.RegisterAppServer(s, &Impl{
		ch: ch,
	})
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	log.Println("Starting server...")
	if err = s.Serve(l); err != nil {
		return
	}

}
