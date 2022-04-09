FROM golang:latest

COPY . ./temp

RUN go mod download
RUN chmod +x rabbitmq.sh && ./rabbitmq.sh

ENV WORKER_COUNT=5
EXPOSE 8080

CMD ["sudo", "rabbitmqctl", "start_app", "&", "&&", "go", "run", "./worker", "&", "&&", "go", "run", "./server"]