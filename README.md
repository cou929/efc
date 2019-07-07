# efc

Error printing format checker.

efc reports printing format which does not use `%+v` directive for error type.

Intend to be applied to projects which uses [pkg/errors](https://github.com/pkg/errors). See also https://godoc.org/github.com/pkg/errors#hdr-Formatted_printing_of_errors

## Insall

```
go get -u github.com/cou929/efc
```

## Example

```go
package main

import "fmt"

func main() {
	var err error
	fmt.Printf("err %v", err)
}
```

efc reports this source like below:

```
$ efc -c=0 ./...
path/to/file/sample.go:7:2: should use %+v format for error type
7               fmt.Printf("err %v", err)
```
