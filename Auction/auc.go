package Auction

import (
	pb "MiniProject3/Auction/proto"
	context "context"
	"fmt"

	"sync"
)

type AuctionItem struct{
	AuctionHighestBid int32
	mu sync.Mutex
}



type Auctionserver struct {
	Id int32
	pb.UnimplementedAuctionServer
}

func (c *Auctionserver) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	
	fmt.Printf("Server Id: %v \n Request from Client: %v", c.Id, in.Id)
	return &pb.BidResponse{Response: "success"}, nil
}

func (c *Auctionserver) Result(ctx context.Context, in *pb.GetResult) (*pb.HighestResult, error) {
	
	c.AddClient(client{clientId: in.Id})
	 //fmt.Printf("Server Id: %v \n Request from Client: %v", c.Id, in.Id)
	return &pb.HighestResult{Result: AucObject.AuctionHighestBid, AucDone: false}, nil
	
}

func (c *Auctionserver) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	
	return &pb.UpdateResponse{Response: "success"}, nil
}

func (c *Auctionserver) AddClient(client client) {

	if len(clientObject.CQue) == 0 {
		clientObject.mu.Lock()
		clientObject.CQue = append(clientObject.CQue, client)
		clientObject.mu.Unlock()
	} else {
		if !Equals(client.clientId) {
			clientObject.mu.Lock()
			clientObject.CQue = append(clientObject.CQue, client)
			clientObject.mu.Unlock()
		}
	}
}

func Equals(i int32) bool {
	for _, v := range clientObject.CQue {
		if v.clientId == i {
			return true
		}
	}
	return false
}

type client struct{
	clientId int32
	highestBid int32
}

type que struct {
	CQue []client
	mu   sync.Mutex
}

var clientObject = que{}
var AucObject = AuctionItem{}