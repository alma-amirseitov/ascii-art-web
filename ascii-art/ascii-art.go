package asciiArt

import (
	"errors"
	"strings"
)

func GetStyled(text string, style string) (string, error) {
	if len(text) == 0 {
		return "", nil
	}

	err := validateText(text)
	if err != nil {
		return "", err
	}

	m, err := InitializeAsciiArtMap(style)
	if err != nil {
		return "", err
	}

	strArr := strings.Split(text, "\n")
	out := ""

	for j, s := range strArr {
		if j == 0 && s == "" && strArr[1] == "" {
			continue
		} else if s == "" {
			out += "\n"
			continue
		}

		for i := 0; i < 8; i++ {
			newStr := ""
			for _, r := range s {
				strArr := strings.Split(m[r], "\n")
				newStr += strArr[i]
			}
			out += newStr + "\n"
		}
	}
	return out, nil
}

func validateText(text string) error {
	for _, r := range text {
		if r == '\n' {
			continue
		} else if r >= ' ' && r <= '~' {
			continue
		} else {
			return errors.New("invalid text")
		}
	}
	return nil
}
