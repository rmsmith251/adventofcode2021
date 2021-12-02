package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	// "strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	
	count := 0
	// prev := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		fmt.Print(current)
		// cur, err := strconv.Atoi(current)
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(2)
		// }
		// if cur > prev && prev != 0 {
		// 	count++
		// }
		// prev := cur
	}
	fmt.Print(count)
}