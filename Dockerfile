FROM rabbitmq:latest

COPY --from=golang:latest /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /temp

COPY . .

RUN go mod download all

ENV WORKER_COUNT=5
EXPOSE 8080

CMD ["rabbitmqctl", "start_app", "&", "&&", "go", "run", "./worker", "&", "&&", "go", "run", "./server"]