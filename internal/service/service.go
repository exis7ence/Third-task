package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutomaticCodeDetection(s string) (string, error) {
	if strings.TrimSpace(s) == "" {
		return "", errors.New("empty input")
	}

	if isMorseCode(s) {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}

func isMorseCode(s string) bool {
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '\n' {
			return false
		}
	}
	return true
}
