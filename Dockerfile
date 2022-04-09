FROM golang:latest

RUN apt install curl
RUN curl -s https://packagecloud.io/install/repositories/rabbitmq/rabbitmq-server/script.deb.sh | sudo bash
RUN curl -s https://packagecloud.io/install/repositories/rabbitmq/erlang/script.deb.sh | sudo bash
RUN apt install rabbitmq-server

WORKDIR /temp

COPY . .

RUN go mod download all

ENV WORKER_COUNT=5
EXPOSE 8080

CMD ["go", "run", "./worker", "&", "&&", "go", "run", "./server"]