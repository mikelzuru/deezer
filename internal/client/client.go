package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mikelzuru/deezer/info"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Search(cfg *Config) error {
	//flag.Parse()
	//c := GetSearcherClient(cfg)
	// Set up a connection to the server.
	conn, err := grpc.Dial(*&cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSearcherClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var r *pb.BasicSearchResponse
	//var err error
	switch {
	case cfg.Type == "artist":
		r, err = c.Search(ctx, &pb.BasicSearchRequest{Query: *&cfg.Query})
	default:
		r, err = c.Search(ctx, &pb.BasicSearchRequest{Query: *&cfg.Query})
	}

	if err != nil {
		log.Fatalf("could not search: %v", err)
	}
	log.Printf("Search response: %s", r.GetResponse())

	return nil
}

func GetSearcherClient(cfg *Config) pb.SearcherClient {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*&cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	return pb.NewSearcherClient(conn)
}
