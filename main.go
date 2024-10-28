package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	port := "8080"
	if len(args) > 0 {
		port = args[0]
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("quotes.txt")
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

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
