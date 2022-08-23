package deezer

import (
	"flag"

	"github.com/glerchundi/subcommands"
	"github.com/mikelzuru/deezer/internal/client"
	"github.com/mikelzuru/deezer/internal/server"
	"github.com/saltosystems/x/log"
)

// NewDeezerCommand create and returns the root cli command
func NewDeezerCommand(logger log.Logger) *subcommands.Command {
	deezerCmd := subcommands.NewCommand("deezer", flag.CommandLine, nil)
	deezerCmd.AddCommand(newServeCommand(logger))
	deezerCmd.AddCommand(newSearchCommand(logger))

	return deezerCmd
}

func newServeCommand(logger log.Logger) *subcommands.Command {
	cfg := &server.Config{}
	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	cfg.ServerPort = fs.Int("port", 50051, "The server port")

	return subcommands.NewCommand(fs.Name(), fs, func() error {
		return server.Create(cfg, logger)
	})
}

func newSearchCommand(logger log.Logger) *subcommands.Command {
	cfg := &client.Config{}
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	fs.StringVar(&cfg.Addr, "addr", "localhost:50051", "the address to connect to")
	fs.StringVar(&cfg.Query, "query", "hasselhoff", "Query to be performed agains music provider")

	return subcommands.NewCommand(fs.Name(), fs, func() error {
		return client.Search(cfg, logger)
	})
}
