# quotes
Quotes as a service

### GET /

> Get a random line of text from the input file.

### POST /

> Add a new quote to the input file. Request body is the new quote.

## Usage
```
-input string
    The file to read lines from (default "quotes.txt")
-port int
    The port to serve from (default 8080)
```
## Build and run
```
go build main.go
./main -port <port> -input <filename>
```

## Build and run with Docker
```
docker build -t quotes .
docker run -v ${PWD}/<filename>:/<filename> -p <port>:<port> quotes -port <port> -input <filename>
```