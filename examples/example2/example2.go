package main

import (
	"fmt"

	"github.com/eliotttak/go-logs"
)

func main() {
	logger := logs.New().
		SetPrefix("[LOG]").
		SetSuffix("- {{{15:04:05}}}")

	s := logger.Slogf("A log of one line. %d + %d = %d", 2, 3, 2+3)
	s += logger.Slog("The first line.\nThe second line.")
	fmt.Print(s)
}
