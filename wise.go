package starterlib

import "github.com/thepudds-example-staging/starterlib/v2/hello"

import "fmt"

func WiseHello(who string) {
        hello.Hello(who)
        fmt.Println("Don't communicate by sharing memory, share memory by communicating.")
}
