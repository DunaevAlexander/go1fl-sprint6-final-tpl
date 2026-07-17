package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func SearchAndConvert (input string) (string, error){
	if input == "" {
		return "", nil
	}

	trimmed := strings.TrimSpace(input)

	isMorse := true
	for _, char := range trimmed {
		if char != '.' && char != '-' && char != ' ' {
			isMorse = false
			break
		}
	}

	if isMorse {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}