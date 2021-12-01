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

	if AucObject.AuctionHighestBid == 0 {
		go timer()
	}

	fmt.Printf("Server Id: %v \n Request from Client: %v \n", c.Id, in.Id)
	if !AuctionDone {
		AucObject.mu.Lock()
		if in.Bid > AucObject.AuctionHighestBid {
			AucObject.AuctionHighestBid = in.Bid
			AucObject.clientId = in.Id
			log.Printf("Bid is successful. \n Bid recieved from: %v. \n Bid is: %v \n", AucObject.clientId, AucObject.AuctionHighestBid)
			mssg := "Your bid has been recieved \n"
			return &pb.BidResponse{Response: "Success", Mssg: mssg, AucDone: AuctionDone}, nil
		} else {
			AucObject.mu.Unlock()
			log.Printf("Bid is unsuccessful. \n Bid recieved from: %v. \n Bid is: %v and needs to be higher than: %v \n", in.Id, in.Bid, AucObject.AuctionHighestBid)
			mssg := "Your bid needs to be higher than:" + strconv.Itoa(int(AucObject.AuctionHighestBid)) + "\n"

			return &pb.BidResponse{Response: "Exception", Mssg: mssg, AucDone: AuctionDone}, nil
		}
	} else {
		log.Printf("The Auction is done. The winner is %v with the highest bid: %v", AucObject.clientId, AucObject.AuctionHighestBid)
		mssg := "The Auction is done. The highest bid was: " + strconv.Itoa(int(AucObject.AuctionHighestBid)) + " and the winner was: " + strconv.Itoa(int(AucObject.clientId)) + "\n"
		return &pb.BidResponse{Response: "Fail", Mssg: mssg, AucDone: AuctionDone}, nil
	}
	//AucObject.AuctionHighestBid

}

func (c *Auctionserver) Result(ctx context.Context, in *pb.GetResult) (*pb.HighestResult, error) {

	//c.AddClient(client{clientId: in.Id})
	//fmt.Printf("Server Id: %v \n Request from Client: %v", c.Id, in.Id)
	log.Printf("The current highest bid is: %v", AucObject.AuctionHighestBid)
	return &pb.HighestResult{Result: AucObject.AuctionHighestBid}, nil
}

func (c *Auctionserver) Done(ctx context.Context, in *pb.DoneRequest) (*pb.Empty, error) {
	AucObject.mu.Unlock()

	return &pb.Empty{}, nil
}

func timer() {

	time.Sleep(30 * time.Second)

	AuctionDone = true

	// Do heavy work
}

//https://yourbasic.org/golang/time-reset-wait-stop-timeout-cancel-interval/

var AucObject = AuctionItem{}
