package main

import (
	Auction "MiniProject3/Auction/proto"
	"fmt"

	"context"
	"log"
	"sync"

	"google.golang.org/grpc"
)

var HighestBid int32
var MyLastBid int32
	
func main() {
	//connect to grpc server
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	conn2, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to gRPC server :: %v", err)
	}
	defer conn.Close()
	defer conn2.Close()

	clientConfig(conn, conn2)


	//message stream
	fmt.Println("Client Node: ", HighestBid)
}

//sets name for client and status
func clientConfig(conn *grpc.ClientConn, conn2 *grpc.ClientConn) {

	//reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Welcome to The Auction! \n")
	fmt.Printf("Enter your id : ")
	var input int
	_, err := fmt.Scan(&input)
	fmt.Printf("Your id is: %v", input)
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}
	var tmp []Auction.AuctionClient
	tmp = append(tmp, Auction.NewAuctionClient(conn), Auction.NewAuctionClient(conn2))
	
	//call ChatService to create a stream
	clientObj := client {clientid: int64(input), aucClients: tmp}
	
	clientObj.Bid()
	//ch.clientName = strings.Trim(name, "\r\n")
	//ch.SendStatus()
}

type client struct{
	clientid int64
	aucClients []Auction.AuctionClient
	mu   sync.Mutex
}

func (client *client) Bid(){
	fmt.Printf("Enter your bid : ")
	var bid int32
	_, err := fmt.Scan(&bid)
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}

	if bid <= HighestBid {
		fmt.Printf("Your bid must be higher than %v", HighestBid)
	} else {
		MyLastBid = bid

		for _, v := range client.aucClients{

			BidResponse, err := v.Bid(context.Background(), &Auction.BidRequest{Id: int32(client.clientid), Bid: MyLastBid})
			
			if err != nil {
				log.Fatalf("Failed to call AuctionService :: %v", err)
			}
			if BidResponse.Response == "success" {
				fmt.Printf("Your bid has been recieved")
				client.Result()
			} else if BidResponse.Response == "The Auction is done"{
				fmt.Printf(BidResponse.Response)
			} 
		}	
	}
}


func (client *client) Result(){

	
	HighestBid, err := client.aucClients[0].Result(context.Background(), &Auction.GetResult{Id: int32(client.clientid), GetResult: "What is the result"})
	HighestBid2, err := client.aucClients[1].Result(context.Background(), &Auction.GetResult{Id: int32(client.clientid), GetResult: "What is the result"})

	if err != nil {
		log.Fatalf("Failed to call AuctionService :: %v", err)
	}
	
	fmt.Printf("The highest bid at the moment is: %v \n", HighestBid)
	fmt.Printf("The highest bid at the moment is: %v \n", HighestBid2)

	client.Bid()
}


