package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	banner := string(data)
	banner = strings.ReplaceAll(banner, "\r", "")
	bannerList := strings.Split(banner, "\n")

	input := os.Args[1]
	SplitInput := strings.Split(input, "\\n")

	for i, word := range SplitInput {
		if word == "" && i == 0 {
			continue
		}
		if word == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {

			for _, char := range word {
				if char < 32 || char > 126 {
					fmt.Println("Character usage out of range")
				}

				startIndex := (int(char) - 32) * 9
				fmt.Print(bannerList[startIndex+i])
			}
			fmt.Println()
		}
	}

}
