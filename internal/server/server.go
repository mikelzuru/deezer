package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/mikelzuru/deezer/info"
	"google.golang.org/grpc"
)

var (
	//port             = flag.Int("port", 50051, "The server port")
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

// Search implements info.Searcher
func (s *server) Search(ctx context.Context, in *pb.BasicSearchRequest) (*pb.BasicSearchResponse, error) {
	log.Printf("Received: %v", in.GetQuery())
	return &pb.BasicSearchResponse{Response: search(in.GetQuery())}, nil
}

// func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
// }

func Create(cfg *Config) error {
	//flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *cfg.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearcherServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func search(q string) string {
	url := musicProviderUrl + "search?q=" + q
	var deez DeezerData

	//Llamamos al proveedor de musica y obtenemos un JSON de respuesta
	getJson(url, &deez)
	fmt.Println("Track n1:", json2string(deez.Data[0]))
	//Transformamos a un JSON con estructura reducida y devolvemos
	return json2string(deez)
}

func json2string(j interface{}) string {
	out, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return string(out)
}

// Variable para la funcion 'getJson'
var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
