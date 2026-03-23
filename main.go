package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run: \"Hello\"")
		return
	}

	data, err := os.ReadFile("standard.txt")

	if err != nil {
		fmt.Println("Error Eeading File", err)
		os.Exit(1)
	}

	banner := string(data)
	//fmt.Print(banner)

	banner = strings.ReplaceAll(banner, "\r", "")

	bannerTwo := strings.Split(banner, "\n")

	input := os.Args[1]

	line := strings.Split(input, "\\n")

	for i, word := range line {
		if word == "" && i == 0 {
			continue
		}

		if word == "" {
			fmt.Println()
			continue
		}

		for i := 1; i <= 8; i++ {
			for _, ascii := range word {
				if ascii < 32 || ascii > 126 {
					continue
				}
				start := (int(ascii) - 32) * 9
				fmt.Print(bannerTwo[start+i])

			}
			fmt.Println()
		}

	}

}
