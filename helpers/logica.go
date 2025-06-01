package helpers

import (
	"strings"
	"unicode"
)

// Contador de vocales (mayúsculas y minúsculas)
func ContarVocales(s string) int {
	vocales := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vocales, char) {
			count++
		}
	}
	return count
}

// Contador de consonantes (solo letras, excluye vocales y símbolos)
func ContarConsonantes(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) && !strings.ContainsRune("aeiouAEIOU", char) {
			count++
		}
	}
	return count
}

// Concatenar iniciales en mayúsculas
func Iniciales(s string) string {
	words := strings.Fields(s)
	var result strings.Builder
	for _, word := range words {
		firstRune := []rune(word)[0]
		result.WriteRune(unicode.ToUpper(firstRune))
	}
	return result.String()
}
