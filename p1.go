package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 1000; i++ {
        if (i % 3) < 1 || (i % 5) < 1 {
            sum = sum + i
        }
    }
    fmt.Printf("%d\n", sum)
}
