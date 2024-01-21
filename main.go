package main

import (
	"fmt"
	"github.com/djotto/rss-transform/config"
	_ "net/http/pprof" //nolint:gosec
	"os"
)

func main() {
	if err := config.Init(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "main::main(): %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s: %+v\n", config.Dir, config.AppConfig)

	for _, pipelineConfig := range config.PipelineConfigs {
		fmt.Printf("%+v\n", pipelineConfig)
	}

}
