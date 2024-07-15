FROM golang:alpine
ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -v -o ./encryption ./cmd/main.go

ENTRYPOINT ["./encryption", "-config", "./configs/dev.yml"]

EXPOSE 8000