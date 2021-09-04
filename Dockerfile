FROM golang:1.16-alpine

RUN mkdir /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /product-grpc-server

EXPOSE 8080

CMD [ "/product-grpc" ]