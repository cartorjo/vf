// sample.go

package main

import "fmt"

func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

func main() {
    personName := "Jose"
    greeting := greet(personName)

    fmt.Println(greeting)
}
