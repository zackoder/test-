package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	if len(args) != 1 {
		fmt.Println("you can't use more or less the tow arguments")
		return
	}
	output := args[0]
	r := []rune(output)
	for i := 0; i < len(r); i++ {
		if r[i] < 32 && r[i] > 177 {
			fmt.Println("you need to choose a character from the ascii table.\nfor more informations about the ascii table google it or run the command man ascii")
			return
		}
	}
	for i := 0; i < len(r); i++ {
		printingchar(r[i], file)
	}
}

func printingchar(r rune, file *os.File) {
	n := int(r - 32)
	n = n + 2 + (n)*8
	startLine := n
	// fmt.Println(startLine)
	endLine := startLine + 8

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Specify the start and end line numbers

	// Scan lines from the file
	lineNumber := 1
	var lines []string
	for scanner.Scan() {
		lineNumber++
		if lineNumber < startLine {
			continue // Skip until the start line
		}
		if lineNumber > endLine {
			break // Stop scanning after the end line
		}

		line := scanner.Text() // Get the current line
		// fmt.Println(line)      // Process the line (print in this case)
		lines = append(lines, line+"\n")
	}
	fmt.Println(lines)

	// Check for errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
