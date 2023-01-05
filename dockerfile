FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /product-data-search

EXPOSE 8080

CMD ["/product-data-search"]
