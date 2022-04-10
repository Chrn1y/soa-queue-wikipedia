FROM golang:latest

WORKDIR /temp

COPY . .

RUN go mod download all

EXPOSE 8080

ENTRYPOINT ["go", "run", "./server"]