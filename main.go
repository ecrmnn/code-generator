package main

import (
	"fmt"
	"math/rand"
	"time"
	"flag"
	"bytes"
	"log"
	"os"
	"strings"
)

var (
	letters = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
	}

	numbers = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}

	codes = map[string]bool{}
)

func main() {
	start := time.Now()

	pattern := flag.String("pattern", "llll-nn-ll-nnnn-lnln", "Set code pattern")
	length := flag.Int("length", 1, "Number of codes to generate")
	letterCase := flag.String("case", "lower", "User upper or lower case letters")
	flag.Parse()

	// Generate codes
	generator(*pattern, *length, *letterCase)

	// Open file for writing and write codes
	file, err := os.Create(currentDirectory() + "/codes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for key, _ := range codes {
		file.WriteString(key + "\n")
	}

	fmt.Println("\nDone. Codes stored in codes.txt")

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func generator(pattern string, length int, letterCase string) {
	for len(codes) < length {
		generateCode(pattern, letterCase)
		fmt.Printf("\rGenerated %d of %d codes", len(codes), length)
	}
}

func generateCode(pattern, letterCase string) {
	buffer := bytes.Buffer{}

	for _, bytePosition := range pattern {
		if (bytePosition == 108) {
			letter := randomCharacter(letters)

			if (letterCase == "upper") {
				letter = strings.ToUpper(letter)
			}

			buffer.WriteString(letter)
		} else if (bytePosition == 110) {
			buffer.WriteString(randomCharacter(numbers))
		} else {
			buffer.WriteString("-")
		}
	}

	codes[buffer.String()] = true
}

func randomCharacter(character []string) string {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))

	for i := len(character) - 1; i > 0; i-- {
		j := rand.Intn(i)
		character[i], character[j] = character[j], character[i]
	}

	return character[0]
}

func currentDirectory() string {
	directory, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return directory;
}