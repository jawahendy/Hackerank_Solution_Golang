package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)

    for i := 0; i < a; i++ {
        for j := 0; j < a-i-1; j++ {
            fmt.Print(" ")
        }
        for j := 0; j <= i; j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }
}