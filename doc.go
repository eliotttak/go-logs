// Copyright © 2026 Eliott Takvorian
//
// This file is a part of go-logs.
//
// go-logs is a free software: you can redistribute it and/or modify it under the terms of
// the GNU Lesser General Public License as published by the Free Software Foundation,
// either version 3 of the License, or (at your option) any later version.
//
// go-logs is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for
// more details.
//
// You should have received a copy of the GNU Lesser General Public License along with
// go-logs. If not, see <https://www.gnu.org/licenses/>.

// Package logs implements a basic logger system. Each log is a
// separate paragraph, and each line has a prefix and a suffix
// (not in the lines separating two paragraphs). The suffix are
// padded for only one log. You can insert the timestamp in the
// content, prefix or suffix of a log by inserting a
// [time.Time.Format]() layout string in triple curly braces
// (e.g. '{{{Mon. Jan. 2 2006 15:04:05}}}'). You can add colors
// with ANSI code colors (see [github.com/eliotttak/go-ansi-colors]).
//
// Example 1 - Basics (usage, prefix, suffix, timestamp):
//
//	package main
//
//	import "github.com/eliotttak/go-logs"
//
//	func main() {
//		logger := logs.New("[LOG]", "- {{{15:04:05}}}")
//		logger.Logf("A log of one line. The number is %d", 3)
//		logger.Log("The first line.\nThe second line.")
//	}
//
//	// Output:
//	//
//	// [LOG] A log of one line. The number is 3 - 12:08:34
//	//
//	// [LOG] The first line.  - 12:08:34
//	// [LOG] The second line. - 12:08:34
//
// Example 2 - In a string:
//
//	package main
//
//	import (
//		"github.com/eliotttak/go-logs"
//		"fmt"
//	)
//
//	func main() {
//		logger := logs.New("[LOG]", "- {{{15:04:05}}}")
//		s := logger.Slogf("A log of one line. The number is %d", 3)
//		s += logger.Slog("The first line.\nThe second line.")
//		fmt.Print(s)
//	}
//
//	// The output is the same as above, but you can see that the log passes by a string before being displayed.
//
// Example 3 - In a buffer (file, STDERR, etc.):
//
//	package main
//
//	import (
//		"log"
//		"os"
//
//		"github.com/eliotttak/go-logs"
//	)
//
//	func main() {
//		file, err := os.Create("example_3.txt")
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer file.Close()
//
//		logger := logs.New("[LOG]", "- {{{15:04:05}}}")
//		logger.Flogf(file, "A log of one line. The number is %d", 3)
//		logger.Flog(file, "The first line.\nThe second line.")
//	}
//
//	// If you open the file example_3.txt after having executed this program, it will contain the same text as in the two examples above.
//
// You can run theses examples by running 'go run github.com/eliotttak/go-logs/examples/exampleX' where X is the number of the example.
package logs
