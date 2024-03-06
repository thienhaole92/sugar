# sugar

`sugar` is a Golang library that leverages generics to prove various methods. It aims to port as many small ultilities from other languages to `Go`. The implementation is highly influenced by `Python`.

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
