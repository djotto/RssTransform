package main

import (
	"fmt"
	"github.com/djotto/rss-transform/config"
	_ "net/http/pprof" //nolint:gosec
	"os"
)

func main() {
	if err := config.Init(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s: %+v\n", config.Dir, config.AppConfig)
}
