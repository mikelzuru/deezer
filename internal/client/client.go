package client

import (
	"context"
	"time"

	pb "github.com/mikelzuru/deezer/info"
	"github.com/saltosystems/x/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Search(cfg *Config, logger log.Logger) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*&cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSearcherClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Search(ctx, &pb.BasicSearchRequest{Query: *&cfg.Query})
	if err != nil {
		logger.Error("could not search: %v", err)
	}
	logger.Info("Search response: %s", r.GetResponse())

	return nil
}
