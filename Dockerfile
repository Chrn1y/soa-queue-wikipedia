FROM golang:latest

WORKDIR /temp

COPY . .

RUN go mod download all
RUN chmod +x rabbitmq.sh && ./rabbitmq.sh

ENV WORKER_COUNT=5
EXPOSE 8080

CMD ["sudo", "rabbitmqctl", "start_app", "&", "&&", "go", "run", "./worker", "&", "&&", "go", "run", "./server"]