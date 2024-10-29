FROM golang:bookworm
WORKDIR /app
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /quotes main.go
WORKDIR /
ENTRYPOINT ["/quotes"]
