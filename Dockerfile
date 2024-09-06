FROM golang:1.22.4 AS dev

WORKDIR /app

COPY . ./

RUN go mod tidy
# RUN go build -o main

EXPOSE 8080

CMD ["go","run","main.go"]