package main

import (
	"fmt"
	"github.com/djotto/rss-transform/config"
	"github.com/djotto/rss-transform/pipeline"
	_ "net/http/pprof" //nolint:gosec
	"os"
	"sync"
	"time"
)

func main() {
	// Load all the config files
	if err := config.Init(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "main::main(): %v\n", err)
		os.Exit(1)
	}

	// Initialize pipelines
	var wg sync.WaitGroup
	cancel, err := pipeline.Init(config.PipelineConfigs, &wg)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "main::main(): Failed to initialize pipelines: %v\n", err)
		os.Exit(1)
	}

	time.Sleep(30 * time.Second)

	cancel()
	wg.Wait()
}
