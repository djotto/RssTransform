// contract.go - This file contains all the exported members of the pipeline
// package intended for public use. Other variables in the package may be
// exported, but are not intended for external consumption.

// Package pipeline handles concurrent ETL pipelines for RSS feeds.
package pipeline

import (
	"context"
	"fmt"
	"github.com/djotto/rss-transform/config"
	"sync"
)

// Init spins up one goroutine (see rssPipeline.go) per pipeline configuration.
func Init(pipelineConfigs []config.PipelineConfigStruct, wg *sync.WaitGroup) (context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())

	rssPipelines := make([]*rssPipeline, 0, len(pipelineConfigs))
	for _, pipelineConfig := range pipelineConfigs {
		wg.Add(1)
		rssPipeline, err := newRssPipeline(pipelineConfig)
		if err != nil {
			return cancel, fmt.Errorf("pipeline::Init(): %w", err)
		}

		rssPipeline.start(ctx, wg)

		rssPipelines = append(rssPipelines, rssPipeline)
	}

	return cancel, nil
}
