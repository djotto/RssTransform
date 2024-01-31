.PHONY: all rss-pipeline

all: rss-pipeline

rss-pipeline:
	go build -o bin/rss-pipeline ./cmd/rss-pipeline
