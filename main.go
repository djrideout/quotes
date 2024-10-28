package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var fileName = flag.String("input", "quotes.txt", "The file to read lines from")
var port = flag.Int("port", 8080, "The port to serve from")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(*fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		line := lines[rand.Intn(len(lines))]
		fmt.Fprint(w, line)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
