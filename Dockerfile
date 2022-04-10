FROM golang:latest

WORKDIR /temp

COPY . .

RUN go mod download all

EXPOSE 8080

CMD ["go", "run", "./server"]