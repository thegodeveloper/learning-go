# Learning Go

Mastering Golang.

## Professional Go Developer

- Learn the basics
- Build a REST API
- Microservices with gRPC
- Build & Deployment
- Build a CLI tool
- Performance
- Testing
- Frontend Apps with Gin Framework
- Learn Hugo
- Deploy applications to Google Cloud Run

## Books

- The Go Programming Language
- Learning Go
- The Power of Go Tools
- Concurrency in Go
- Distributed Services with Go
- Network Programming with Go
- Go for DevOps
- 100 Go Mistakes
- Building Distributed Applications in Gin
- Build Systems with Go
- Mastering Go
- Let's Go
- Let's Go Further
- Microservices with Go
- Building Modern CLI Applications with Go

## Install goimports

Enhanced version of `go fmt` called `goimports`.

```shell
go install golang.org/x/tools/cmd/goimports@latest
```

To run it across a project:

```shell
goimports -l -w .
```

- The `-l` flat tells `goimports` to print the files with incorrect formatting to the console.
- The `-w` flag tells `goimports` to modify the files in-place.
- The `.` specifies the files to be scanned: everything in the current directory and all of its subdirectories.
