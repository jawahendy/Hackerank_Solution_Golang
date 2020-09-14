package main

import (
    "fmt"
    "math"
)

func main() {

    var n int
    fmt.Scan(&n)

    a := make([][]int, n)

    leftsum, rightsum := 0, 0

    for i := 0; i < n; i++ {
        a[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&a[i][j])
            if i == j {
                leftsum += a[i][j]
            }
            if i+j == n-1 {
                rightsum += a[i][j]
            }
        }
    }

    different := math.Abs(float64(leftsum - rightsum))
    fmt.Println(different)

}