# Builder stage
FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Runner stage
FROM scratch as runner
COPY --from=builder /app/myapp /myapp
ENTRYPOINT ["/myapp"]

# Development stage
FROM golang:1.21 as dev
WORKDIR /app
COPY go.mod ./
COPY go.sum* ./
RUN go mod download
COPY . .
CMD ["go", "run", "./..."]
