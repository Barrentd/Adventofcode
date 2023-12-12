package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

// Global map for number words to their digit equivalents
var numberWords = map[string]string{
    "one":   "o1e",
    "two":   "t2o",
    "three": "t3e",
    "four":  "f4r",
    "five":  "f5e",
    "six":   "s6x",
    "seven": "s7n",
    "eight": "e8t",
    "nine":  "n9e",
}

func replaceLetterNumbers(input string) string {
    //fmt.Println("inp:", input)	
    result := input
    for word, digit := range numberWords {
        result = strings.ReplaceAll(result, word, digit)
    }
    //fmt.Println("out:", result)
    return result
}

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

    reDigit := regexp.MustCompile(`\d`)

    for scanner.Scan() {
        line := scanner.Text()
	//fmt.Println("Linewithletter:", line)
	result := replaceLetterNumbers(line)
	//fmt.Println("Linewithdigit:", result)
	
	numbers := reDigit.FindAllString(result, -1)
	combinedNumberStr := string(numbers[0]) + string(numbers[len(numbers)-1])
        combinedNumber, _ := strconv.Atoi(combinedNumberStr)
        //fmt.Printf("Numbers combined: %s\n", combinedNumberStr)

        sum += combinedNumber

    }
    fmt.Printf("Total Sum: %d\n", sum)

}

