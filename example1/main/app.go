package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Top10("cat and dog, one dog,two cats and one man"))
}

func Top10(srcString string) []string {
	words := strings.Fields(srcString)
	wordCounter := make(map[string]int)

	for _, word := range words {
		wordCounter[word]++
	}

	type entrySlice struct {
		key   string
		value int
	}

	var sortedSlice []entrySlice

	for key, value := range wordCounter {
		sortedSlice = append(sortedSlice, entrySlice{key, value})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[j].value < sortedSlice[i].value
	})

	var resSlice []string
	for i := 0; i < min(10, len(sortedSlice)); i++ {
		resSlice = append(resSlice, sortedSlice[i].key)
	}

	return resSlice
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(srcString string) (string, error) {
	var builder strings.Builder

	if srcString == "" {
		return "", nil
	}

	if isCorrect, err := Verify(srcString); isCorrect && err == nil {
		runes := []rune(srcString)
		size := len(runes)
		for i := 0; i < size-1; i++ {
			if !unicode.IsDigit(runes[i+1]) && !unicode.IsDigit(runes[i]) {
				builder.WriteRune(runes[i])
			}
			if !unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+1]) {
				repeatCount, _ := strconv.Atoi(string(runes[i+1]))
				if repeatCount != 0 {
					builder.WriteString(strings.Repeat(string(runes[i]), repeatCount))
				}
			}
		}
		if !unicode.IsDigit(runes[size-1]) {
			builder.WriteRune(runes[size-1])
		}
	} else {
		return "", err
	}

	return builder.String(), nil
}

func Verify(srcString string) (bool, error) {

	runes := []rune(srcString)
	if unicode.IsDigit(runes[0]) {
		return false, ErrInvalidString
	}

	for i := 1; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i-1]) {
			return false, ErrInvalidString
		}
	}

	return true, nil
}
