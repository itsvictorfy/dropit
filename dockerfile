FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY templates/* ./templates/
COPY templates/css/* ./templates/css/
COPY templates/js/* ./templates/js/

RUN go build -o /dropIt

EXPOSE 8080

CMD ["/dropIt"]
