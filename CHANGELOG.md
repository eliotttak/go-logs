# `go-logs` changelog

 ---

## v1.0.0 - ![First version][first]
- Support for prefix and suffix
- Support for timestamp inserting
- Support for writing in a [`string`](https://pkg.go.dev/builtin#string) and an [`io.Writer`]
- `Log()`, `Logf()`, `Slog()`, etc. functions similar to
  [`fmt.Print()`](https://pkg.go.dev/fmt#Print), [`fmt.Print()`](https://pkg.go.dev/fmt#Printf),
  [`fmt.Print()`](https://pkg.go.dev/fmt#Sprint), etc.

 ---

## v2.0.0 ![Major version][2.0.0] ![Latest version][latest]
- New features:
	+ Adding of support for default writer, to permise automatically logging to a specific writer (that can be any [`io.writer`])
  	+ New (and better) way to create a logger (see the examples)
- Bug corrections:
	+ Correct issues with CRLF logs

 ---

 [first]: https://img.shields.io/badge/v1.0.0-First_version-blue "This is the first published version"
 [2.0.0]: https://img.shields.io/badge/v2.0.0-Major_version-red "This is a major version"
 [latest]: https://img.shields.io/badge/v2.0.0-Latest_version-01bf01 "This is the latest version published"
 [`io.writer`]: https://pkg.go.dev/io#Writer