package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Print("\n\n" + strconv.Itoa(r.Intn(10)+1) + "\n\n")
}
