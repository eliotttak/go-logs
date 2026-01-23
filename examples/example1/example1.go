package main

import "github.com/eliotttak/go-logs"

func main() {
	logger := logs.New("[LOG]", "- {{{15:04:05}}}")
	logger.Logf("A log of one line. The number is %d", 3)
	logger.Log("The first line.\nThe second line.")
}
