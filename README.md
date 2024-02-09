# Notes for Go

## Installation

Install `go` 1.22.0 in `/usr/local`

```bash
curl -OL https://golang.org/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.22.0.linux-amd64.tar.gz
```

Run the following line or add them to `.bashrc`

```bash
export PATH=/usr/local/go/bin:$PATH
```

Initialize the `go.mod`

```bash
go mod init example.com/hello
```

Update dependencies

```bash
go mod tidy
```

To build your go code

```bash
go build
```

You can look up modules here <https://pkg.go.dev>.

## Module

- Function with capital letter at the start is an exported function!

- `:=` is define and assignment 2 in 1

- for local dependency, can do a simple hack fix like the following

  ```bash
  go mod edit -replace example.com/greetings=../greetings
  ```

  The command specifies that `example.com/greetings` should be replaced with `../greetings` for the purpose of locating the dependency.

  The run this to synchronize the dependency

  ```bash
  go mod tidy
  ```
  
- `for` is Go's `while`

- Download external dependencies

  ```bash
  go get golang.org/x/example/hello/reverse
  ```

## Workspace

- Initialize workspace

  ```bash
  go work init ./hello
  ```

- Add additional module to the workspace

  ```bash
  go work use ./example/hello
  ```

## Testing

- Name the test file with `_testing.go`

- All functions should start with `func Test..()`

- Run the test using

  ```bash
  go test -v
  ```

## Install

Check the install path

```bash
go list -f '{{.Target}}'
```

Change bin location

```bash
go env -w GOBIN=/path/to/your/bin
```

Reset the bin location

```bash
go env -u GOBIN
```

Install the binary

```bash
go install
```

## Cheat sheet

- Effective Go - <https://go.dev/doc/effective_go>
- Go module references - <https://go.dev/ref/mod>

