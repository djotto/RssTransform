# Builder stage
FROM golang:1.21 as build
WORKDIR /app
COPY go.mod ./
COPY go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rss-transform .

# Runner stage
FROM scratch as run
COPY --from=build /app/rss-transform /rss-transform
ENTRYPOINT ["/rss-transform"]

# Development stage
FROM golang:1.21 as dev
WORKDIR /app
COPY go.mod ./
COPY go.sum* ./
RUN go mod download
COPY . .
CMD ["go", "run", "./..."]
