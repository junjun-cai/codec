# codec
this is a encode/decode tool by golang, include base2 to base64 etc.

## Features
* multiple methods to encode/decode.
* customize encode/decode std str.
* test/benchmark every codec.

## installation
```bash
go get -u github.com/junjun-cai/codec
```

## Getting Started
```go
package main

import (
	"fmt"
	"github.com/junjun-cai/codec/base32"
)

func main() {
	ret, _ := base32.StdCodec.Encode([]byte("hello world"))
	fmt.Println(ret)
}
// Output: [78 66 83 87 89 51 68 80 69 66 51 87 54 52 84 77 77 81 61 61 61 61 61 61]
```