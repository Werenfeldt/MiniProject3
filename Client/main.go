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

}

//sets name for client and status
func clientConfig(conn *grpc.ClientConn, conn2 *grpc.ClientConn) {

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
	clientObj := client{clientid: int64(input), aucClients: tmp}

	for _, v := range clientObj.aucClients {
		fmt.Println("clients", v)

	}

	clientObj.Bid()

}

type client struct {
	clientid   int64
	aucClients []Auction.AuctionClient
	mu         sync.Mutex
}

func (client *client) Bid() {
	fmt.Printf("Enter your bid : ")
	var bid int32
	_, err := fmt.Scan(&bid)
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}

	MyLastBid = bid

	var BidResponse [2]*Auction.BidResponse

	for i := 0; i < len(client.aucClients); i++ {

		BidResponse[i], err = client.aucClients[i].Bid(context.Background(), &Auction.BidRequest{Id: int32(client.clientid), Bid: MyLastBid})

		if err != nil {
			client.remove(i)
		}

	}

	for i := 0; i < len(client.aucClients); i++ {

		if BidResponse[i] != nil {
			if BidResponse[i].Response == "success" {
				fmt.Printf("Your bid has been recieved \n")
				client.Result()
			} else if BidResponse[i].Response == "Your bid needs to be higher than:" {
				fmt.Printf(BidResponse[i].Response, "%v", BidResponse[i].Timestamp)
				client.Bid()
			} else if BidResponse[i].Response == "The Auction is done" {
				fmt.Printf(BidResponse[i].Response)
			}
		}
	}

}

func (client *client) Result() {

	var HighestBidtmp *Auction.HighestResult

	for i := 0; i < len(client.aucClients); i++ {
		var err error

		HighestBidtmp, err = client.aucClients[i].Result(context.Background(), &Auction.GetResult{Id: int32(client.clientid), GetResult: "What is the result"})

		if err != nil {
			log.Fatalf("Failed to call AuctionService :: %v", err)

		}
	}

	fmt.Printf("The highest bid at the moment is: %v \n", HighestBidtmp.Result)

	HighestBid = HighestBidtmp.Result

	for i := 0; i < len(client.aucClients); i++ {
		var err error

		Done, err := client.aucClients[i].Update(context.Background(), &Auction.UpdateRequest{ClientId: int32(client.clientid), AucDone: false})
		if err != nil {
			log.Fatalf("Failed to call AuctionService :: %v, %v", err, Done)

		}
	}

	client.Bid()
}

func (client *client) remove(i int) {
	client.aucClients[i] = client.aucClients[len(client.aucClients)-1]
	client.aucClients = client.aucClients[:len(client.aucClients)-1]
	return
}
