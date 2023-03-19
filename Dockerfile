# Build stage
FROM golang:1.19-alpine3.16 as dev

WORKDIR /app

COPY . ./

CMD [ "go","run","main.go" ]
