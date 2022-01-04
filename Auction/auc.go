package Auction

import (
	pb "MiniProject3/Auction/proto"
	context "context"
	
	"log"
	"strconv"
	"sync"
	"time"
)

//Rename 
type Auctionserver struct {
	pb.UnimplementedAuctionServer
}


//Struct to save highest bid and who made it. 
type AuctionItem struct {
	AuctionHighestBid int32
	clientId          int32
	mu                sync.Mutex
}

var AucObject = AuctionItem{}

//The timer for the auction to be done
var AuctionDone bool

func (c *Auctionserver) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	
	if AucObject.AuctionHighestBid == 0 {
		go timer()
	}
	
	AucObject.mu.Lock()
	if !AuctionDone {
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
		AucObject.mu.Unlock()
		log.Printf("The Auction is done. The winner is %v with the highest bid: %v", AucObject.clientId, AucObject.AuctionHighestBid)
		mssg := "The Auction is done. The highest bid was: " + strconv.Itoa(int(AucObject.AuctionHighestBid)) + " and the winner was: " + strconv.Itoa(int(AucObject.clientId)) + "\n"
		return &pb.BidResponse{Response: "Fail", Mssg: mssg, AucDone: AuctionDone}, nil
	}


}

func (c *Auctionserver) Result(ctx context.Context, in *pb.GetResult) (*pb.HighestResult, error) {

	log.Printf("The current highest bid is: %v", AucObject.AuctionHighestBid)
	return &pb.HighestResult{Result: AucObject.AuctionHighestBid}, nil
}

func (c *Auctionserver) Done(ctx context.Context, in *pb.DoneRequest) (*pb.Empty, error) {
	AucObject.mu.Unlock()

	return &pb.Empty{}, nil
}

func timer() {
	time.Sleep(120 * time.Second)
	AuctionDone = true
}


