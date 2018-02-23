package utils

import (
	"math/rand"
	"time"
)

func HumanReadablePassword(alphabetSize, numberSize int) string {
	rand.Seed(time.Now().UnixNano())
	vowels := "aeiou"
	consonants := "bcdfghjklmnpqrstvwxyz"
	digits := "0123456789"

	prefixSize := alphabetSize / 2
	if alphabetSize%2 != 0 {
		prefixSize = int(alphabetSize/2) + 1
	}
	suffixSize := alphabetSize - prefixSize

	var prefixPart = make([]byte, prefixSize)

	for i := 0; i <= prefixSize-1; i++ {
		if i%2 == 0 {
			// use consonants
			prefixPart[i] = consonants[rand.Intn(len(consonants)-1)]
		} else {
			// use vowels
			prefixPart[i] = vowels[rand.Intn(len(vowels)-1)]
		}
	}

	var midPart = make([]byte, numberSize)

	// use digits
	for k := range midPart {
		midPart[k] = digits[rand.Intn(len(digits))]
	}

	var suffixPart = make([]byte, suffixSize)

	for i := 0; i <= suffixSize-1; i++ {
		if i%2 == 0 {
			// use consonants
			suffixPart[i] = consonants[rand.Intn(len(consonants)-1)]
		} else {
			// use vowels
			suffixPart[i] = vowels[rand.Intn(len(vowels)-1)]
		}
	}

	return string(prefixPart) + string(midPart) + string(suffixPart)
}

