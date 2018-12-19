package starterlib

import "github.com/thepudds-example-staging/starterlib/hello"

import "fmt"

func WiseHello() {
        hello.Hello("World")
        fmt.Println("Don't communicate by sharing memory, share memory by communicating.")
}
