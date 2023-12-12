package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run script.go filename")
        os.Exit(1)
    }

    filename := os.Args[1]
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("Error opening file: %s\n", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    sum := 0

    // Regular expression to find individual digits
    re := regexp.MustCompile(`\d`)

    for scanner.Scan() {
        line := scanner.Text()
        //fmt.Println("Line:", line)

        // Extract digits into a slice
        digits := re.FindAllString(line, -1)

	//fmt.Println("digits:", digits)
	numbers := digits[0]+digits[len(digits) -1]

	//fmt.Println("Number", numbers)

	result, _ := strconv.Atoi(numbers)
	sum += result
        }

    fmt.Printf("Total Sum: %d\n", sum)
}

