package go-hello

import "fmt"

func Hello() {
    fmt.Println("Hello, world!")
}

func Sum(nums ...int) int {
    sum := 0
    for _, val := range nums {
        sum += val
    }

    return sum
}

