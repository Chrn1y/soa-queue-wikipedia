FROM golang:latest

WORKDIR /wikipedia-server

COPY ../. .

RUN go mod download

EXPOSE 8080
EXPOSE 5672

CMD ["go", "run", "server/."]