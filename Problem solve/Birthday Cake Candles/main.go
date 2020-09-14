package main

import "fmt"

func main() {
    var b, c, max, occ int
    fmt.Scan(&b)

    a := []int{}
    for i := 0; i < b; i++ {
        fmt.Scan(&c)
        a = append(a, c)
        if c > max {
            max = c
        }
    }

    for _, num := range a {
        if num == max {
            occ++
        }
    }

    fmt.Println(occ)

}