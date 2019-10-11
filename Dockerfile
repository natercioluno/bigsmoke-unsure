FROM golang:1.12

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR player
RUN go build -o main player/main.go

EXPOSE 9000-10000

ENTRYPOINT ["./main"]
