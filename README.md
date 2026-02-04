# `go-logs`

`go-logs` is a Go package that implements a logging system to improve logging in e.g. a web server.

It is designed to be simple to use, while offering many features.

## Main features:
- Log in many places, including support for `stdout`, any `io.Writer` as `stderr` or a file, and
  even a string.
- A prefix and a suffix for each line of log
- Easy inserting of the timestamp for each log
- Setting of a default writer to automate writing in, for instance,  a specific file

More information in the API documentation following.

## Documentation

See the `go-doc` documentation [on Go Packages](https://pkg.go.dev/github.com/eliotttak/go-logs).
If the link is invalid, you can find doc in the source files.

## Changelog

See [CHANGELOG.md](./CHANGELOG.md).

## Licensing

Copyright &copy; 2026 Eliott TAKVORIAN

This program is published under the GNU Lesser General Public License version 3. You can use this
code in any project, even if your project is proprietary, but if you modify `go-logs`, you have to
publish the modifications under the GNU Lesser General Public License version 3, or, at your option,
any later version. You don't have to credit me, although I would appreciate it as it would improve
my visibility on GitHub and elsewhere.

This is not a legal advice; you can find a full copy of the GNU Lesser General Public License version 3 in the file
[LICENSE.md](./LICENSE.md).
