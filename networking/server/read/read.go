package read

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// ReadTextfile returns the content of the given textfile formated as CLI input
func ReadTextfile(filename string) string {
	var result string
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	var split []byte = ([1]byte{'n'})[0:1]
	lines := bytes.Split(dat, split)

	for _, line := range lines {
		for i, letter := range line {
			if letter == ':' {
				result += string(line[:i])
				result += " "
			}
		}
	}
}
