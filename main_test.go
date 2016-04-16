package main

import (
	"os"
	"bufio"
	"testing"
)

func TestCanGenerateOneHundredCodesConsistingOfFourLettersInFile(t *testing.T) {
	codes := []string{}

	runMainWithFlags([]string{"-pattern=llll", "-length=100"}, &codes)

	if len(codes) != 100 || len(codes[0]) != 4 {
		t.Error("Could not generate file with 100 codes with length of 4")
	}
}

func runMainWithFlags(flags []string, arrayToPopulate *[]string) {
	os.Args = append(os.Args, flags...)
	main()

	readFileAndPopulateArray(arrayToPopulate)
}

func readFileAndPopulateArray(array *[]string) {
	file, err := os.Open(currentDirectory() + "/codes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*array = append(*array, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func TestInvalidPatternWillResultInDashes(t *testing.T)  {
	invalidPattern := "Awesome Sauce"
	codes = map[string]bool{}

	generator(invalidPattern, 1, "lower")

	if codes["-------------"] != true {
		t.Error("Invalid pattern did not result in dashes")
	}
}