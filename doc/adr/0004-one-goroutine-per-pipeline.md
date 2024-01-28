# 4. One goroutine per pipeline

Date: 2024-01-28

## Status

Accepted

## Context

Each pipeline is expected to spend long periods asleep (see ACR 0002). There
are two basic approaches we can take: Have each process manage its own timer,
or have a single management process that spins up goroutines when it's time for
a pipeline to run.

## Decision

I decided to have one goroutine per pipeline, with each pipeline managing its
own timer.

Alternatives considered: An off-the-shelf threadpool implementation with a
single goroutine managing all the timers.

The *Recovery* and *Resource utilization* risks, below, are strong enough
arguments that I might revisit this decision in a future version, but for now
development speed tips the balance against premature optimization.

## Consequences

Advantages:

1. *Simplicity*: Each pipeline managing its own timer simplifies the
   architecture, making the system easier to implement and reason about.
2. *Isolation*: Once the pipelines are running, there's no single point of
   failure.
3. *Development speed*: Avoiding the threadpool package means one less thing
   to learn.

Risks:

1. *Recovery*: In a threadpool architecture a failed pipeline would naturally
   just retry in a few minutes. We're going to need to think about how
   pipelines should recover from failure.
2. *Resource utilization*: Goroutines are light, but they still consume
   resources. This solution is likely to have higher resource utilization when
   dealing with large numbers of pipelines simultaneously. In a threadpool
   implementation, the maximum number of threads can be limited.
