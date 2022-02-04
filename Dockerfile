# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
COPY . /mygrpc
WORKDIR /mygrpc
EXPOSE 50051

CMD [ "go", "run", "server/main.go" ]