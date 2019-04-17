package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	fileName *string
)

func init() {
	fileName = flag.String("file", "", "File full path")
	flag.Parse()

	if fileName == nil {
		panic("fileName is null")
	}
	if len(*fileName) == 0 {
		panic("fileName is null")
	}
}

func main() {
	lineCount := 0
	file, err := os.Open(*fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Printf("%v: %v\n", lineCount, scanner.Text())
		lineCount++
	}
}
