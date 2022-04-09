FROM golang:latest

WORKDIR /wikipedia-worker

COPY ../. .

RUN go mod download

EXPOSE 5672

ENV WORKER_COUNT=5

CMD ["go", "run", "worker/."]