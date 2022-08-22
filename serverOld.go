// Package main implements a server for Deezer Search service.
package old

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/mikelzuru/deezer_grpc/info"
	"google.golang.org/grpc"
)

var (
	port             = flag.Int("port", 50051, "The server port")
	musicProviderUrl = "https://api.deezer.com/"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedSearcherServer
}

type DeezerData struct {
	Data []struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Duration int    `json:"duration"`
		Artist   struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"artist"`
		Album struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"album"`
		Type string `json:"type"`
	} `json:"data"`
	Total int    `json:"total"`
	Next  string `json:"next"`
}

// SayHello implements info.GreeterServer
func (s *server) Search(ctx context.Context, in *pb.BasicSearchRequest) (*pb.BasicSearchResponse, error) {
	log.Printf("Received: %v", in.GetQuery())
	return &pb.BasicSearchResponse{Response: search(in.GetQuery())}, nil
}

// func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
// }

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearcherServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func search(q string) string {
	url := musicProviderUrl + "search?q=" + q
	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// return readJsonFromBody(resp)

	//Si queremos transformarlo a un JSON con estructura reducida
	var deez DeezerData
	getJson(url, &deez)
	fmt.Println("Track n1:", json2string(deez.Data[0]))
	return json2string(deez)
}

func json2string(j interface{}) string {
	out, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func readJsonFromBody(resp *http.Response) string {
	resultJson := ""

	for {
		bs := make([]byte, 99999)
		n, err := resp.Body.Read(bs)
		resultJson += string(bs[:n])
		if err != nil {
			break
		}
	}

	return resultJson
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
