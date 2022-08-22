package main

import (
	"flag"
	"os"
	"strings"

	"github.com/mikelzuru/deezer/internal/cli/deezer"
	"github.com/peterbourgon/ff"
	sdlog "github.com/saltosystems/x/log/stackdriver"
)

func main() {
	// create logger
	logger := sdlog.New()
	deezerCmd := deezer.NewDeezerCommand(logger)

	err := deezerCmd.Execute(os.Args[1:], func(fs *flag.FlagSet, args []string) error {
		return ff.Parse(fs, args,
			ff.WithConfigFileFlag("config"),
			ff.WithConfigFileParser(ff.PlainParser),
			ff.WithEnvVarPrefix(strings.ToUpper(deezerCmd.Name())),
		)
	})
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
