package main

import (
	"fmt"

	"github.com/eliotttak/go-logs"
)

func main() {
	logger := logs.New("[LOG]", "- {{{15:04:05}}}")
	s := logger.Slogf("A log of one line. The number is %d", 3)
	s += logger.Slog("The first line.\nThe second line.")
	fmt.Print(s)
}
