package asciiArt

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"os"
	"strings"
)

var (
	TEMPLATE_STANDARD_HASH_BYTE   = []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}
	TEMPLATE_SHADOW_HASH_BYTE     = []byte{38, 185, 77, 11, 19, 75, 119, 233, 253, 35, 224, 54, 11, 253, 129, 116, 15, 128, 251, 127, 101, 65, 209, 216, 197, 216, 94, 115, 238, 85, 15, 115}
	TEMPLATE_THINKERTOY_HASH_BYTE = []byte{243, 125, 219, 114, 118, 66, 122, 15, 188, 217, 227, 105, 214, 89, 103, 10, 252, 146, 207, 178, 243, 230, 71, 217, 19, 204, 220, 6, 169, 152, 222, 67}
)

// FileClose function to close the text file
func FileClose(file *os.File) error {
	err := file.Close()
	if err != nil {
		return err
	}
	return nil
}

// InitializeAsciiArtMap function of making maps between rune and printing string
func InitializeAsciiArtMap(templateName string) (map[rune]string, error) {
	//-----------------------------------------
	templateNames := [3]string{"shadow", "standard", "thinkertoy"}
	if !IsStringInStrArray(templateName, templateNames) {
		return nil, errors.New("invalid banner type, avilable : standard, shadow, thinkertoy")
	}

	templateFileName := "./ascii-art/styles/" + strings.ReplaceAll(templateName, " ", "") + ".txt"

	readFile, err := os.Open(templateFileName)
	defer FileClose(readFile)
	if err != nil {
		return nil, err
	}

	validFile := CheckFileValidity(templateFileName)
	if !validFile {
		return nil, errors.New("empty file or Corrupted file to open")
	}

	//-----------------------------------------
	scanner := bufio.NewScanner(readFile)
	scanner.Scan()

	m := make(map[rune]string) // newMap that stores runes' ascii-art
	var r rune = ' '
	i := 1
	for scanner.Scan() { // Every eight lines are stored in map as ascii-art
		if i == 9 { // with corresponding key rune
			i = 1
			r++ // key rune is updated every eight lines
			continue
		}
		if i != 8 {
			m[r] += scanner.Text() + "\n"
		} else {
			m[r] += scanner.Text()
		}

		i++
	}
	return m, nil
}

// CheckRunes checks the inporper runes that do not have ascii-art
func CheckRunes(str string) bool {
	for _, r := range str {
		if r >= ' ' && r <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}

func IsStringInStrArray(str string, list [3]string) bool {
	for _, e := range list {
		if e == str {
			return true
		}
	}
	return false
}

func CheckFileValidity(templateFileName string) bool {
	banner, _ := os.ReadFile(templateFileName)
	h := sha256.New()
	h.Write(banner)

	switch templateFileName {
	case "./src/standard.txt":
		if string(h.Sum(nil)) != string(TEMPLATE_STANDARD_HASH_BYTE) {
			return false
		}
	case "./src/shadow.txt":
		if string(h.Sum(nil)) != string(TEMPLATE_SHADOW_HASH_BYTE) {
			return false
		}
	case "./src/thinkertoy.txt":
		if string(h.Sum(nil)) != string(TEMPLATE_THINKERTOY_HASH_BYTE) {
			return false
		}
	}

	return true
}
