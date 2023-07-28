FROM golang:latest

LABEL maintainer="Luca Maggio <lucamaggio1992@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV PORT 8000

RUN go build -o main .

CMD ["./main"]