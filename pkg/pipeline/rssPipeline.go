package pipeline

import (
	"context"
	"fmt"
	"github.com/djotto/rss-transform/pkg/config"
	"sync"
	"time"
)

// rssPipeline is a struct that encapsulates a single ETL pipeline for processing RSS feeds.
type rssPipeline struct {
	id     int
	config config.PipelineConfigStruct
}

var idCounter int

// newRssPipeline creates and initializes a new RssPipeline instance.
func newRssPipeline(cfg config.PipelineConfigStruct) (*rssPipeline, error) {
	// Create a new instance of RSSPipeline with the provided configuration.
	id := func() int {
		idCounter++
		return idCounter
	}()

	fmt.Println("Instantiating goroutine", id)

	return &rssPipeline{
		config: cfg,
		id:     id,
	}, nil
}

// start runs the rssPipeline's main logic in a goroutine.
func (rp *rssPipeline) start(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("Starting goroutine", rp.id)

	duration := time.Duration(rp.config.SleepDuration) * time.Second
	timer := time.NewTimer(duration)

	go func() {
		defer wg.Done()
		defer timer.Stop()
		for {
			// Implement the main logic of the pipeline here.
			select {
			case <-ctx.Done():
				// Perform cleanup and exit goroutine
				rp.shutdown()
				return
			case <-timer.C:
				fmt.Printf("Config: %+v\n", rp.config)
				fmt.Println("Firing timer in goroutine", rp.id)
				timer.Reset(duration)
			}
		}
	}()
}

// shutdown gracefully terminates the pipeline. It releases resources,
// closes any open connections, and performs any necessary cleanup.
func (rp *rssPipeline) shutdown() {
	fmt.Println("Shutting down goroutine", rp.id)
	// Implement shutdown logic
}
