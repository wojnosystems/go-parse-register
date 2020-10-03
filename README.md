# Overview

Parse Register stores functions that converts string representations of values into Go structures.

`go get github.com/wojnosystems/go-parse-register`

```go
package main
import (
  "fmt"
  "github.com/wojnosystems/go-parse-register"
  "reflect"
)
type myType struct {
  Name string
}
func main() {
  reg := &parse_register.Registry{}
  reg.Register(reflect.TypeOf((*myType)(nil)).Elem(), func(settableDst interface{}, value string) (err error) {
    out := settableDst.(*myType)
    out.Name = value
    return nil
  } )
  
  foo := myType{}
  _, _ = reg.SetValue(&foo, "bar")
  fmt.Println(foo.Name)
}
```

The above program will output "bar". It creates a new registry, registers a custom handler. This handler is kept simple for this example, but you can do much more with it. 
