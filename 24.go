package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := x
    for i := 0; i < 1000; i++ {
        prevZ := z
        z = z - (z*z - x)/(2*z);
        if math.Abs(prevZ-z) < 1e-10 {
//            fmt.Println(i)
            return z
        }
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
