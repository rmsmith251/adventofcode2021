package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	
	depth := 0
	horizontal := 0
	aim := 0
	scanner := bufio.NewScanner(file)
	// i := 0
	for scanner.Scan() {
		current := scanner.Text()
		line := strings.Split(current, " ")
		change, _ := strconv.Atoi(line[1])
		if line[0] == "forward" {
			if aim != 0 {
				depth += (aim * change)
			}
			horizontal += change
		} else if line[0] == "down" {
			aim += change
		} else if line[0] == "up" {
			aim -= change
		} else{
			fmt.Println("Error")
		}
		// fmt.Println("depth:", depth, "horizontal:", horizontal, "aim:", aim, "depth*horiz:", depth * horizontal)
		// i++
		// if i > 5 {
		// 	os.Exit(1)
		// }
	}
	fmt.Println(depth, horizontal, aim, depth * horizontal)
}