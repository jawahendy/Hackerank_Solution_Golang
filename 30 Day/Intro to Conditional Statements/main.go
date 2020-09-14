package main

import "fmt"

var a int

func main() {
    fmt.Scan(&a)
    if a%2 != 0 {
        fmt.Println("Weird")
    }
    if a%2 == 0 {
        for i:=2; i<6; i++ {
            if a == i {
                fmt.Println("Not Weird")
            }
        }
    }
    if a%2 == 0 {
        for i:=6; i<21; i++ {
            if a == i {
                fmt.Println("Weird")
            }
        }
    }
    if a%2 == 0 && a > 20 {
        fmt.Println("Not Weird")
    }
}