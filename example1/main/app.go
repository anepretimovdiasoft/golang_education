package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Unpack("44"))
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
