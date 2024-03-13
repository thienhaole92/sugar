# sugar

`sugar` is a Golang library that leverages generics to support aggregate & transform functions in Go. It aims to port as many small ultilities from other languages to `Go`. The implementation is highly influenced by `Python` and `JavaScript`.


## Requirement

- Go 1.20

## Installation

```sh
go get github.com/thienhaole92/sugar
```

## Usage

1. Import the package.

```go
import "github.com/thienhaole92/sugar"
```

2. Call a function. For example:

```go
result := sugar.Map([]int{1, 2, 3, 4}, func(k int) int {
	return k - 1
})

fmt.Printf("%v", result)
//output: [0 1 2 3]
```
