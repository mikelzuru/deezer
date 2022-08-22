package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	pb "github.com/mikelzuru/deezer/info"
	"github.com/saltosystems/x/log"
	"google.golang.org/grpc"
)

var (
	musicProviderUrl = "https://api.deezer.com/"
)

// server is used to implement info.SearcherServer.
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
	return &pb.BasicSearchResponse{Response: search(in.GetQuery())}, nil
}

func Create(cfg *Config, logger log.Logger) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *cfg.ServerPort))
	if err != nil {
		logger.Error("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearcherServer(s, &server{})
	logger.Info("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Error("failed to serve: %v", err)
	}

	return nil
}

func search(q string) string {
	url := musicProviderUrl + "search?q=" + q
	var deez DeezerData

	//Llamamos al proveedor de musica y obtenemos un JSON de respuesta
	getJson(url, &deez)
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
