package main

import (
	"fmt"
	"os"

	"github.com/eliotttak/go-logs"
)

func main() {
	logger := logs.New().
		SetPrefix("[ERROR]").
		SetSuffix("- {{{15:04:05}}}").
		SetDefaultWriter(os.Stderr)

	logger.Log("CRITICAL ERROR")
	logger.Log("Something very bad occured:\n writing /bin/foobar.txt: permission denied")
	fmt.Print("Note: this message is false; the program didn't try to write to /bin/foobar.txt\n\n")
}
