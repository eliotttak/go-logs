package main

import "github.com/eliotttak/go-logs"

func main() {
	logger := logs.New().
		SetPrefix("[LOG]").
		SetSuffix("- {{{15:04:05}}}")

	logger.Logf("A log of one line. %d + %d = %d", 2, 3, 2+3)
	logger.Log("The first line.\nThe second line.")
}
