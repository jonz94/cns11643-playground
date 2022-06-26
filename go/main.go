package main

import (
	"fmt"
)

func main() {
	words := "123abc全字庫。𧃽，󿷴不是貓"

	for _, r := range words {
		fmt.Printf("%c (%U) ", r, r)
		if r >= 0xF0000 {
			fmt.Printf("=> plus\n")
		} else if r >= 0x20000 {
			fmt.Printf("=> ext-b\n")
		} else {
			fmt.Printf("=> normal\n")
		}
	}
}
