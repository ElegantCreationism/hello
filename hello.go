package main

import "fmt"

var names = []string{"mohamed","nathan", "alex", "john"}


func main() {

    for _, name := range names {
        fmt.Printf("%v, welcome back to your page.\n", CustomGreet(name))
    }

}

func Greet() string {
    return "Hello There!!"
}

func CustomGreet(name string) string {
    return "Hello " + name
}
