package main

import "fmt"

func main() {
    var hour, minute, second int
    var suffix string
    fmt.Scanf("%d:%d:%d%s", &hour, &minute, &second, &suffix)
    if suffix == "AM" && hour == 12 {
        hour = 0
    }
    if suffix == "PM" && hour != 12 {
        hour += 12
    }
    fmt.Printf("%02d:%02d:%02d", hour, minute, second)
}