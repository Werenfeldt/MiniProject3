package main

import (
	"fmt"
	"log"
	"net"
	"os"

	Auc "MiniProject3/Auction"
	proto "MiniProject3/Auction/proto"

	"google.golang.org/grpc"
)

func main() {
	//Init log file
	
	LOG_FILE := "../Auction_log"

	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	//log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	fmt.Printf("Welcome to The Auction! \n")
	fmt.Printf("Enter your id : ")
	var input int32
	_, err = fmt.Scan(&input)
	
	fmt.Printf("Your id is: %v", input)
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}


	//Init port
	Port := os.Getenv("PORT")
	if Port == "" {
		if input == 1 {

			Port = "8080" //default Port set to 5000 if PORT is not set in env
		} else {
			Port = "5000" //default Port set to 5000 if PORT is not set in env
		}
	}



	//Init listener
	listen, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", Port, err)
	}
	log.Println("Listening @ : " + Port)
	fmt.Println("Listening @ : " + Port)

	//gRPC server instance
	grpcserver := grpc.NewServer()

	//register Service
	cs := Auc.Auctionserver{Id: input}
	proto.RegisterAuctionServer(grpcserver, &cs)

	//grpc listen and serve
	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	} else {

	}
}