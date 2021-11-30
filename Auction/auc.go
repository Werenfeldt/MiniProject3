package Auction

import (
	pb "MiniProject3/Auction/proto"
	context "context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

type AuctionItem struct {
	AuctionHighestBid int32
	clientId          int32
	mu                sync.Mutex
}

type Auctionserver struct {
	Id int32
	pb.UnimplementedAuctionServer
}

var AuctionDone bool

func (c *Auctionserver) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {

	go timer()

	fmt.Printf("Server Id: %v \n Request from Client: %v \n", c.Id, in.Id)
	if !AuctionDone {
		AucObject.mu.Lock()
		if in.Bid > AucObject.AuctionHighestBid {
			AucObject.AuctionHighestBid = in.Bid
			AucObject.clientId = in.Id
			return &pb.BidResponse{Response: "success"}, nil
		} else {
			AucObject.mu.Unlock()
			mssg := "Your bid needs to be higher than:" + strconv.Itoa(int(AucObject.AuctionHighestBid))

			return &pb.BidResponse{Response: mssg}, nil
		}
	} else {
		mssg := "The Auction is done. The highest bid was: " + strconv.Itoa(int(AucObject.AuctionHighestBid)) + " and the winner was: " + strconv.Itoa(int(AucObject.clientId))
		return &pb.BidResponse{Response: mssg}, nil
	}
	//AucObject.AuctionHighestBid

}

func (c *Auctionserver) Result(ctx context.Context, in *pb.GetResult) (*pb.HighestResult, error) {

	//c.AddClient(client{clientId: in.Id})
	//fmt.Printf("Server Id: %v \n Request from Client: %v", c.Id, in.Id)

	return &pb.HighestResult{Result: AucObject.AuctionHighestBid}, nil
}

func (c *Auctionserver) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	AucObject.mu.Unlock()

	return &pb.UpdateResponse{Response: "success"}, nil
}

func timer() {
	timer := time.AfterFunc(time.Minute, func() {
		log.Println("Foo run for more than a minute.")
		AuctionDone = true
	})

	defer timer.Stop()

	// Do heavy work
}

//https://yourbasic.org/golang/time-reset-wait-stop-timeout-cancel-interval/

var AucObject = AuctionItem{}
