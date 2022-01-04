package main

import (
	//"fmt"
	"log"
	"net"

	Auc "MiniProject3/Auction"
	proto "MiniProject3/Auction/proto"

	//"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	//Old writing logs. 
	//Init log file
	//LOG_FILE := "../Auction_log"
	
	// logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer logFile.Close()
	
	// log.SetOutput(logFile)


	//Init port
	ownPort := ":5000"
	

	log.Printf("Listening on port %v", ownPort)

	//Init listener
	listen, err := net.Listen("tcp", ownPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %v, %v", ownPort, err)
	}
	
	//gRPC server instance
	grpcserver := grpc.NewServer()

	//register Service
	cs := &Auc.Auctionserver{}
	proto.RegisterAuctionServer(grpcserver, cs)

	//grpc listen and serve
	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	} 

	//Old crash register

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// go func(){
	// 	for sig := range c {
	// 		log.Printf("The Server: %v has been: %v. The server is dead.", sig)
	// 		os.Exit(1)
	// 		// sig is a ^C, handle it
	// 	}
	// }()
}

