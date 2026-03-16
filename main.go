package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) >3 {
		fmt.Println("incorrect Input: use: go run . \"Hello\" The Font")
		return
	}

	input := os.Args[1]

	fontName := "standard.txt"

	if len(os.Args) == 3 {
		fontName = os.Args[2] + ".txt"
	}

	font, err := loadFonts(fontName)
	if err != nil {
		fmt.Println("Failed to load font:", err)
		return
	}
	fmt.Println("Font length:", len(font))
	render(input, font)
}
func render(text string, font []string) {

	if text == "" {
		return
	}

	lines := strings.Split(text, "\n")

	for i, line := range lines {

		if line != "" {
			checkLine(line, font)
		}

		// print newline between blocks only
		if i < len(lines)-1 {
			fmt.Println()
		}
	}
}
func loadFonts(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := strings.ReplaceAll(string(data), "\r", "")
	lines := strings.Split(content, "\n")

	return lines, nil
}

func checkLine(text string, font []string) {

	height := 9

	for row := 0; row < height; row++ {

		for _, char := range text {

			ascii := int(char)

			if ascii < 32 || ascii > 126 {
				continue
			}

			offset := (ascii - 32) * height

			fmt.Print(font[offset+row])
		}

		fmt.Println()
	}
}
