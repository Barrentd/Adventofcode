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
    //bag := []int{12r, 13g, 14b}

    // Regular expression to find individual digits
    regame:= regexp.MustCompile(`^.*:`)
    regrp := regexp.MustCompile(`[^;]+;`)
    renum := regexp.MustCompile(`\d+\w`)
    //re := regexp.MustCompile(`\d\w`)
    delspc := regexp.MustCompile(`\s+`)
    r := regexp.MustCompile(`\d+`)
    n := regexp.MustCompile(`[a-z]`)

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println("Line:", line)

	game := regame.FindString(line)
	gamenum := r.FindString(game)

	rmvgame := regame.ReplaceAllString(line, " ")

        grpnum := regrp.FindAllString(rmvgame, -1)
	// Boucle pour supprimer les espaces de chaque chaîne
	for i, texte := range grpnum {
		grpnum[i] = delspc.ReplaceAllString(texte, "")
	}

	fmt.Printf("grpnum:%d %s\n", len(grpnum), grpnum)
	
	addnum := true
	for _, num := range grpnum {
		numbers := renum.FindAllString(num, -1)
		fmt.Println("Numberis", numbers)
		//num := re.FindAllString(numbers, -1)
		//fmt.Println("Number", num)

		for _, obj := range numbers{
			if n.FindString(obj) == "b" {
				if n, _ := strconv.Atoi(r.FindString(obj)) > 14 {
					addnum = false
                                }
			}
			if  n.FindString(obj) == "r" {
				if n, _ := strconv.Atoi(r.FindString(obj)) > 12 {
					addnum = false
                                }
			}
			if  n.FindString(obj) == "g" {
				if n, _ := strconv.Atoi(r.FindString(obj)) > 13 {
					addnum = false
				}
			}	
		}

        }
	if addnum == true{
		sum += gamenum
	}
	fmt.Printf("Game:%s\n", gamenum)
    }

    fmt.Printf("Total Sum: %d\n", sum)
}
