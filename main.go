package main

import (
    "fmt"
)

func main() { 
    fmt.Println("How many dice do you want?")  
    dice := NewDiceList(ScanInteger())
    dice.Roll()
}