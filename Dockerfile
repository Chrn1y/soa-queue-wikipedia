FROM docker:latest

COPY . .

CMD docker build -t server server/Dockerfile
CMD docker build -t worker worker/Dockerfile

CMD ["docker", "run", "-p", "5672:5672", "-d", "bitnami/rabbitmq:latest", "&&",
     "docker", "run", "-p", "5672:5672", "-p", "8080:8080", "-d", "server", "&&",
     "docker", "run", "-p", "5672:5672", "-d", "worker"]