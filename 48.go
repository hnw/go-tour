package main

import (
    "fmt"
    "math/cmplx"
)
    

func Cbrt(x complex128) complex128 {
    delta := cmplx.Abs(x) * 1e-11
    z := x
    prevZ := complex128(0)
    for cmplx.Abs(prevZ-z) > delta {
        prevZ = z
        z = z - (z*z*z-x)/(3*z*z)
    }
    return z
}

func main() {
    fmt.Println(Cbrt(2))
    fmt.Println(cmplx.Pow(Cbrt(2),3))
}
