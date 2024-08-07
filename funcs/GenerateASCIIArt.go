package funcs

import "strings"

func GenerateASCIIArt(text string, banner []string) string {
	var result strings.Builder
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, r := range line {
				// Ensure the character is within the valid ASCII range
				if r < 32 || r > 126 {
					// Or any other placeholder
					continue
				}
				index := 9*(int(r)-32) + i + 1
				result.WriteString(banner[index])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
