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
var AucDone bool

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

//
//sets name for client and status
func clientConfig(conn *grpc.ClientConn, conn2 *grpc.ClientConn) {

	fmt.Printf("Welcome to The Auction! \n")
	fmt.Printf("Enter your id : ")
	var input int
	_, err := fmt.Scan(&input)
	fmt.Printf("Your id is: %v \n", input)
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}
	var tmp []Auction.AuctionClient
	tmp = append(tmp, Auction.NewAuctionClient(conn), Auction.NewAuctionClient(conn2))

	//call ChatService to create a stream
	clientObj := client{clientid: int64(input), aucClients: tmp}

	clientObj.Bid()

}

type client struct {
	clientid   int64
	aucClients []Auction.AuctionClient
	bidresponse bidresponse
	mu         sync.Mutex
}

type bidresponse struct{
	response []*Auction.BidResponse
	mu         sync.Mutex
}

var BidResponse = bidresponse{response: make([]*Auction.BidResponse, 2)} 

func (client *client) Bid() {
	fmt.Printf("Enter your bid : ")
	var bid int32
	_, err := fmt.Scan(&bid)
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}

	MyLastBid = bid

	var deadServer int
	var anydeadServers bool

	for i := 0; i < len(client.aucClients); i++ {

		BidResponse.response[i], err = client.aucClients[i].Bid(context.Background(), &Auction.BidRequest{Id: int32(client.clientid), Bid: MyLastBid})
		
		if err != nil {
			deadServer = i
			anydeadServers = true
		}
	}

	if anydeadServers {
		BidResponse.mu.Lock()
		client.mu.Lock()
		client.remove(deadServer)
	}

	if BidResponse.response[0].AucDone{
		AucDone = true
	}
	
	if AucDone {
		BidResponse.mu.Lock()
		AucDonefunc(BidResponse.response[0])
	} else {
		for i := 0; i < len(client.aucClients); i++ {
			if BidResponse.response[i] != nil {
				BidResponse.mu.Lock()
				if BidResponse.response[i].Response == "Success" {
					fmt.Printf(BidResponse.response[i].Mssg)
					client.Result()
				} else if BidResponse.response[i].Response == "Exception" {
					fmt.Printf(BidResponse.response[i].Mssg)
					BidResponse.mu.Unlock()
					client.Bid()
				} else {
					BidResponse.mu.Unlock()
				}
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
			log.Fatalf("Result: Failed to call AuctionService :: %v", err)

		}

		err = nil
	}

	fmt.Printf("The highest bid at the moment is: %v \n", HighestBidtmp.Result)

	HighestBid = HighestBidtmp.Result

	client.Done()

}

func (client *client) Done() {

	for i := 0; i < len(client.aucClients); i++ {
	
		Done, err := client.aucClients[i].Done(context.Background(), &Auction.DoneRequest{ClientId: int32(client.clientid)})
		if err != nil {
			log.Printf("Server is down", err, Done)
	
		}

		err = nil
	}

	BidResponse.mu.Unlock()
	
	client.Bid()
}

func (client *client) remove(i int) {
	client.aucClients[i] = client.aucClients[len(client.aucClients)-1]
	BidResponse.response[i] = BidResponse.response[len(BidResponse.response)-1]
	client.aucClients = client.aucClients[:len(client.aucClients)-1]
	BidResponse.response = BidResponse.response[:len(BidResponse.response)-1]

	client.mu.Unlock()
	BidResponse.mu.Unlock()
	return
}

func AucDonefunc(response *Auction.BidResponse){
	fmt.Printf(response.Mssg)
	BidResponse.mu.Unlock()
}
