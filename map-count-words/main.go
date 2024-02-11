package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("great-gatsby.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	book := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		book = append(book, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	onlyWords := []string{}
	for _, v := range book {
		onlyWords = append(onlyWords, strings.Fields(v)...)
	}
	wordCount := make(map[string]int)
	for _, v := range onlyWords {
		wordCount[v] += 1
	}

	for k, v := range wordCount {
		fmt.Printf("%s \t %v\n", k, v)
	}
}
