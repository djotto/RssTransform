# 2. Use Go to implement rss-pipeline

Date: 2024-01-28

## Status

Accepted

## Context

I envisage the `rss-pipeline` component as a long-running service, managing
multiple ETL pipelines that poll data sources periodically and update RSS feeds
if those data sources have changed.

To this end I need a programming language that can efficiently manage
concurrent tasks in a server environment. It's crucial that pipelines do not
block each other, but individual pipelines are inherently linear processes.

I'm not overly concerned about performance, because my intended use case is
less than fifty pipeline processes, each running at most once every thirty
minutes.

## Decision

I decided to use Go (Golang) for implementing this component.

Alternatives considered: Python, PHP, Javascript, Rust.

## Consequences

Advantages:

1. *Learning opportunity*: An opportunity to gain practical experience with
   a small Go project.
2. *Implementation speed*: Go is close enough to other C-descended languages
   that I can hit the ground running. Development shouldn't get bogged
   down with learning.
3. *Concurrency*: Goroutines fulfill my requirement to run multiple
   pipelines simultaneously.
4. *Standard library*: Go's Standard library contains almost all the packages
   I think I'll need.
5. *Service-Oriented Architecture*: Go's design is conducive to building robust
   services in a piecemeal fashion.

Risks:

1. *Learning curve*: I could be underestimating the ease of picking up Go.
2. *Community and resources*: Smaller footprint than, say, Javascript.