package read

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Textfile returns the content of the given textfile formated as CLI input
func Textfile(filename string) string {
	var result string
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	var split = make([]byte, 1)
	split[0] = '\n'
	lines := bytes.Split(dat, split)

	for _, line := range lines {
		for i, letter := range line {
			if letter == ':' {
				result += string(line[:i])
				result += " "
			}
		}
	}
	return result
}
