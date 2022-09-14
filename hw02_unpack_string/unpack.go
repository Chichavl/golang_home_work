package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	ErrInvalidChar   = errors.New("invalid char")
)

func Unpack(s string) (string, error) {
	// convert string to rune array
	runeCodeArr := []rune(s)

	strBuilder := strings.Builder{}

	for pos, runeCode := range runeCodeArr {
		switch {
		// if single digit
		case unicode.IsDigit(runeCode) && pos != 0 && !unicode.IsDigit(runeCodeArr[pos-1]):
			{
				num, err := strconv.Atoi(string(runeCode))
				if err != nil {
					return "", ErrInvalidChar
				}
				// repeat only if digit not 0
				if num != 0 {
					strBuilder.WriteString(strings.Repeat(string(runeCodeArr[pos-1]), num-1))
				} else {
					// if digit is 0 - remove last added char
					str := strBuilder.String()
					if len(str) > 0 {
						temp := []rune(str)
						str = string(temp[:len(temp)-1])
						strBuilder.Reset()
						strBuilder.WriteString(str)
					}
				}
			}
		case !unicode.IsDigit(runeCode):
			// if char - add it to result array
			strBuilder.WriteString(string(runeCode))
		default:
			// if no match trigger invalid string error
			return "", ErrInvalidString
		}
	}

	return strBuilder.String(), nil
}
