package main

import (
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.:50051")
	handleError(err)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError((err))
}
