package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
)

func main() {   
    rand.Seed(time.Now().UnixNano())
    fmt.Print("\n\n" + strconv.Itoa(rand.Intn(10) + 1) + "\n\n")
}
