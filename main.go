package main

import (
	"os"
	"strconv"
	"strings"
)

const alphabet = "ABCDEFGHIJ"

func getDigits(num int) []byte {
	var (
		digits []byte
	)

	s := strconv.Itoa(num)

	for _, n := range s {
		n = n - '0'
		digits = append(digits, alphabet[n])
	}
	return digits
}

func main() {
	var words_file, err = os.ReadFile("words.txt")
	if err != nil {
		panic(err)
	}
	var words = strings.Split(string(words_file), "\n")
	var out = make(map[string]string)
	var out2 = ""
	for i, val := range words {
		out[val] = string(getDigits(i))
		out2 += val + " " + out[val] + "\n"
	}
	os.WriteFile("shorthand.txt", []byte(out2), 0700)
}
