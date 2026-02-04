package main

import (
	"log"
	"os"

	"github.com/eliotttak/go-logs"
)

func main() {
	file, err := os.Create("example_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := logs.New().
		SetPrefix("[LOG]").
		SetSuffix("- {{{15:04:05}}}")

	logger.Log("This log will be written to the standard output.")
	logger.Log("This one also.")
	logger.Flog(file, "This one will be written to './example_3.txt'.")
}
