package utils

import (
	"fmt"
	"os"
)

func ShowLoading(label string, percent int) {

	if percent < 0 || percent > 100 {
		println("percent must be between 0 to 100")
		os.Exit(1)
	}

	if percent == 0 {
		fmt.Print("\033[200D\033[K")
	} else {
		fmt.Print("\033[2A\033[200D\033[K")
	}

	fmt.Println(label)

	//fmt.Print("\033[H\033[2J\033[?25l")
	for i := 0; i < 100; i++ {
		if i < percent {
			fmt.Print("#")
		} else {
			fmt.Print("_")
		}
	}
	fmt.Printf("\t %%%d\n", percent)
}
