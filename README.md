
# Extended JSON support for go

This library is a drop-in replacement for golang's `encoding/json` (actually is a fork of it) that adds support for the "required" fields that is missing from upstream and it seems it's not going to happen soon: https://github.com/golang/go/issues/19858

## Install

Import `go.bug.st/json` instead of `encoding/json`.

## Usage example

```
package main

import (
	"fmt"

	"go.bug.st/json"
)

type TestStruct struct {
	Test string `json:"test,required"`
}

func main() {
	jsondata := []byte(`{ "other":"4" }`)
	var t TestStruct

	fmt.Println(json.Unmarshal(jsondata, &t))

	// Output:
	// json: undefined required field test into main.TestStruct
}
```

https://play.golang.org/p/aCw8d5COcWv

All other JSON operations are 100% compatible with `encoding/json`.

## Development

At the moment the library is a fork of `encoding/json` at version 1.15.6, this library is tagged following the same versioning.

## License

The license is the same from the Golang source code: https://github.com/golang/go/blob/master/LICENSE

## Disclaimer

I do not guarantee that this library is updated regularly from upstream.

