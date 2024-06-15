# Learning Go

This is my personal repository. Where I take notes about my Go journey.

## Starting Year

I started with Go in May 2020.

### June 2024

- 4 years later.
- To learn Go is my priority.

## Books to study in 2024

- Learning Go Second Edition
- Microservices with Go
- gRPC Go for Professionals
- Building Modern CLI Applications in Go
- Event-Driven Architecture in Golang
- Domain-Driven Design in Golang
- Let's Go
- Let's Go Further

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

## go vet

There is another class of errors that developers run into. The code is syntactically valid, but there are mistakes that are not what you meant to do.
This includes things like passing the wrong number of parameters to formatting methods or assigning values to variables that are never used.

Run `go vet` on your project with the command:

```shell
go vet ./...

# github.com/thegodeveloper/learning-go/wTime
wTime/master.go:13:25: 2006-02-01 should be 2006-01-02
```

