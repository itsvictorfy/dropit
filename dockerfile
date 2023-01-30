FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go .
COPY pages/ ./pages

RUN go build -o /dropIt

EXPOSE 8080

CMD ["/dropIt"]
