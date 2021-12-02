package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
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
	prev := 0
	scanner := bufio.NewScanner(file)
	// i := 0
	for scanner.Scan() {
		current := scanner.Text()
		cur, _ := strconv.Atoi(current)
		// fmt.Println(cur, count, prev)
		if cur > prev && prev != 0 {
			count++
		}
		prev = cur
		// i++
		// if i > 5 {
		// 	os.Exit(1)
		// }
	}
	fmt.Println(count)
}