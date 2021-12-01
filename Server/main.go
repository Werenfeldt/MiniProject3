package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

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
	
	fmt.Printf("Welcome to The Auction! \n")
	fmt.Printf("Enter your id : ")
	var input int32
	_, err = fmt.Scan(&input)
	
	fmt.Printf("Your id is: %v \n", input)
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for sig := range c {
			log.Printf("The Server: %v has been: %v. The server is dead.", input, sig)
			os.Exit(1)
			// sig is a ^C, handle it
		}
	}()

	//Init port
	Port := os.Getenv("PORT")
	if Port == "" {
		if input == 1 {

			Port = "8080" 
		} else {
			Port = "5000" 
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