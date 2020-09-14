package main

import "fmt"

func main() {
    var a, b, c, d, t int
    fmt.Scan(&a)

    for i := 0; i < a; i++ {
        fmt.Scan(&t)
        if t > 0 {
            b++
        } else if t < 0 {
            c++
        } else {
            d++
        }
    }

    bf := float64(b) / float64(a)
    cf := float64(c) / float64(a)
    df := float64(d) / float64(a)

    fmt.Printf("%.6f\n", bf)
    fmt.Printf("%.6f\n", cf)
    fmt.Printf("%.6f\n", df)

}