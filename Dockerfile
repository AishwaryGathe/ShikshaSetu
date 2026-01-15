FROM golang:1.22 as base 

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

