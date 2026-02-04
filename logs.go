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

//go:build go1.24

package logs

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var dateHourRegexp = regexp.MustCompile(`\{\{\{.*?\}\}\}`)

// AddPrefixToEachLine adds the specified prefix to each line of s. It puts a space between the
// prefix and the line of text. If prefix is empty, it returns s.
func AddPrefixToEachLine(s string, prefix string) string {
	if prefix == "" {
		return s
	}
	var newString strings.Builder

	splittedStrings := strings.SplitSeq(s, "\n")

	isFirst := true
	for line := range splittedStrings {
		if isFirst {
			isFirst = false
		} else {
			newString.WriteString("\n")
		}
		newString.WriteString(prefix + " " + line)
	}

	return newString.String()
}

// AddSuffixToEachLine adds the specified suffix to each line of s. It pads the suffix to make it
// easily readable. It puts a space between the line of text and the suffix. If suffix is empty, it
// returns s.
func AddSuffixToEachLine(s string, suffix string) string {
	if suffix == "" {
		return s
	}

	s = strings.ReplaceAll(s, "\r\n", "\n")

	// Step 1: split the string to make a list of lines
	splittedStrings := strings.SplitSeq(s, "\n")

	// Step 2: find the length of the longer one
	maxLength := 0
	for line := range splittedStrings {
		lineLength := utf8.RuneCountInString(line)
		maxLength = max(maxLength, lineLength)
	}

	// Step 3: create back the iterator
	splittedStrings = strings.SplitSeq(s, "\n")

	// Step 4: pad and put the suffix
	var newString strings.Builder

	isFirst := true
	for line := range splittedStrings {
		if isFirst {
			isFirst = false
		} else {
			newString.WriteString("\n")
		}
		lineLength := utf8.RuneCountInString(line)
		newString.WriteString(line + strings.Repeat(" ", maxLength-lineLength+1) + suffix)
	}

	return newString.String()
}

// ParseLog creates a log string from a content string, a prefix and a suffix.
//
// You can insert the timestamp by inserting a [time.Time.Format]() layout string in triple curly
// braces (e.g. '{{{Mon. Jan. 2 2006 15:04:05}}}') in the content, the prefix or the suffix.
//
// It adds two new lines at the end to make a space between two logs.
func ParseLog(s string, prefix string, suffix string) string {
	prefixedStr := AddPrefixToEachLine(s, prefix)
	prefixedSuffixedStr := AddSuffixToEachLine(prefixedStr, suffix)

	var result strings.Builder
	result.WriteString(prefixedSuffixedStr)

	for dateHourRegexp.MatchString(result.String()) {
		substr := dateHourRegexp.FindString(result.String())

		layout := strings.ReplaceAll(strings.ReplaceAll(substr, "{", ""), "}", "")

		formattedTime := time.Now().Local().Format(layout)

		temp := result.String()
		result.Reset()
		result.WriteString(strings.Replace(temp, substr, formattedTime, 1))
	}
	result.WriteString("\n\n")

	return result.String()
}

// Logger is a struct that represents one logger system. You should use a pointer and not the direct
// type.
type Logger struct {
	// The prefix of each log. E.g.: "[ERROR]"
	prefix string

	// The suffix of each log. E.g.: "{{{Mon. Jan. 2 2006 15:04:05}}}"
	suffix string

	defaultWriter io.Writer
}

// SetDefaultWriter sets the default writer used by [Logger.Log]() and [Logger.Logf]() to w. It can
// be changed at any time by calling back this method. If you don't define a default writer, the
// default writer will be [os.Stdout]. It returns the logger to permise
// chaining.
func (l *Logger) SetDefaultWriter(w io.Writer) *Logger {
	l.defaultWriter = w
	return l
}

// SetPrefix sets the prefix of each line of log to prefix. It can be changed at any time by calling
// back this method. If you don't define a prefix, it will be "". It returns the logger to permise
// chaining.
func (l *Logger) SetPrefix(prefix string) *Logger {
	l.prefix = prefix
	return l
}

// SetSuffix sets the suffix of each line of log to suffix. It can be changed at any time by calling
// back this method. If you don't define a suffix, it will be "". It returns the logger to permise
// chaining.
func (l *Logger) SetSuffix(suffix string) *Logger {
	l.suffix = suffix
	return l
}

// GetDefaultWriter returns the current default writer of the logger.
func (l *Logger) GetDefaultWriter() io.Writer {
	return l.defaultWriter
}

// GetDefaultWriter returns the current prefix of the logger.
func (l *Logger) GetPrefix() string {
	return l.prefix
}

// GetDefaultWriter returns the current suffix of the logger.
func (l *Logger) GetSuffix() string {
	return l.suffix
}

// Slog works in 3 steps:
//   - Recompose the arguments into a single string
//   - Parse it to a log string (with prefix and suffix)
//   - Return this log string.
//
// For more info about how are the arguments converted into a string, see [fmt.Sprint]().
//
// For more info about how is the string parsed into a log, see [ParseLog]().
func (l *Logger) Slog(a ...any) string {
	return ParseLog(fmt.Sprint(a...), l.prefix, l.suffix)
}

// Slogf is similar to [Logger.Slog]() but instead of just concatenate the arguments, it takes a
// format string. It uses the verbs of [fmt.Print]().
//
// For more info about how are the arguments converted into a string, see [fmt.Sprintf]().
//
// For more info about how is the string parsed into a log, see [ParseLog]().
func (l *Logger) Slogf(format string, a ...any) string {
	return l.Slog(fmt.Sprintf(format, a...))
}

// Flog is similar to [Logger.Slog]() but writes the output into w. It returns the same values as
// [fmt.Fprint]().
//
// For more info about how are the arguments parsed into a log string, see [Logger.Slog]().
//
// For more info about how is this string written into w, see [fmt.Fprint]().
func (l *Logger) Flog(w io.Writer, a ...any) (int, error) {
	return fmt.Fprint(w, l.Slog(a...))
}

// Flog is similar to [Logger.Slogf]() but writes the output into w. It returns the same values as
// [fmt.Fprintf]().
//
// For more info about how are the arguments parsed into a log string, see [Logger.Slogf]().
//
// For more info about how is this string written into w, see [fmt.Fprint]().
func (l *Logger) Flogf(w io.Writer, format string, a ...any) (int, error) {
	return fmt.Fprint(w, l.Slogf(format, a...))
}

// Log is similar to [Logger.Flog]() used with the default writer.
//
// For more info about how the arguments are displayed, see [Logger.Flog]().
func (l *Logger) Log(a ...any) (int, error) {
	return l.Flog(l.defaultWriter, a...)
}

// Logf is similar to [Logger.Flogf]() used with the default writer.
//
// For more info about how the arguments are displayed, see [Logger.Flogf]().
func (l *Logger) Logf(format string, a ...any) (int, error) {
	return l.Flogf(l.defaultWriter, format, a...)
}

// New creates and returns a pointer to a new [Logger] instance.
func New() *Logger {
	return &Logger{
		prefix:        "",
		suffix:        "",
		defaultWriter: os.Stdout,
	}
}
