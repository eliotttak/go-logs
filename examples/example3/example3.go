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

	logger := logs.New("[LOG]", "- {{{15:04:05}}}")
	logger.Flogf(file, "A log of one line. The number is %d", 3)
	logger.Flog(file, "The first line.\nThe second line.")
}
