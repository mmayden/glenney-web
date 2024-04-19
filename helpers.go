package main

import (
	"fmt"
	"strconv"
)

// Get a user input string and convert to integer. Inform user if invalid entry or not an integer.
func ScanInteger() int {
	for {
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			fmt.Println("Invalid entry")
		}
		
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Not an integer.")
		}else {
			if (i==1) {
				fmt.Println("Any number but 1 pal..")
			}else {
			return i
			}
		}
	}
}