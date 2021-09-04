FROM golang:1.16-alpine

RUN mkdir /app

WORKDIR /app

RUN go build -o /product-grpc-server

EXPOSE 8080

CMD [ "/product-grpc" ]