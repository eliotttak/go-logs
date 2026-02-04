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

// Package logs implements a basic logging system. Each log is a separate paragraph, and each line
// can have a prefix and a suffix (not in the lines separating two paragraphs). The suffix are
// padded for only one log. You can insert the timestamp in the content, prefix or suffix of a log
// by inserting a [time.Time.Format]() layout string in triple curly braces (e.g.
// '{{{Mon. Jan. 2 2006 15:04:05}}}'). You can change the default writer of the output, the prefix
// and the suffix.
//
// ========== Example 1 ==========
//
// This example shows you the basics of `go-logs`:
//
//   - Create a logger, with a prefix and a suffix
//   - Logging with [Logger.Log]() and [Logger.Logf]()
//
// ----- Link to the file -----
//
// https://github.com/eliotttak/go-logs/blob/main/examples/example1/example1.go
//
// ----- Run -----
//
//	go run github.com/eliotttak/go-logs/examples/example1@latest
//
// Run it online at https://go.dev/play/p/1TyDL0V84jJ
//
// ========== Example 2 ==========
//
// This example shows you how to log in a string using [Logger.Slog]
// and [Logger.Slogf].
//
// ----- Link to the file -----
//
// https://github.com/eliotttak/go-logs/blob/main/examples/example3/example2.go
//
// ----- Run -----
//
//	go run github.com/eliotttak/go-logs/examples/example2@latest
//
// Run it online at https://go.dev/play/p/6L22CscJaNQ
//
// ========== Example 3 ==========
//
// This example shows you how to temporarily redirect the log with [Logger.Flog]() and
// [Logger.Flogf]().
//
// ----- Link to the file -----
//
// https://github.com/eliotttak/go-logs/blob/main/examples/example3/example3.go
//
// ----- Run -----
//
//	go run github.com/eliotttak/go-logs/examples/example3@latest
//
// Note: since we can't work with files in Go Playground, I didn't create a link for this example.
//
// ========== Example 4 ==========
//
// This example shows how to set a default writer to the logger using [Logger.SetDefaultWriter]().
//
// ----- Link to the file -----
//
// https://github.com/eliotttak/go-logs/blob/main/examples/example4/example4.go
//
// ----- Run -----
//
//	go run github.com/eliotttak/go-logs/examples/example4@latest
//
// Run it online at https://go.dev/play/p/Mn4FnyAMZgE
package logs
