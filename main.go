package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
		file, err := os.OpenFile(*fileName, os.O_RDWR|os.O_APPEND, 666)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		if r.Method == "POST" {
			line, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("Error reading request body:", err)
				return
			}
			defer r.Body.Close()

			lineStr := string(line)
			fmt.Fprintf(file, "\n%s", lineStr)
			fmt.Fprintf(w, "Added \"%s\"", lineStr)
		} else {
			var lines []string
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			line := ""
			if len(lines) > 0 {
				line = lines[rand.Intn(len(lines))]
			}
			fmt.Fprint(w, line)
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
