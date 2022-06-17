package ascii

import (
	"fmt"
	"strings"
)

const FontHeight = 8
const AsciiStart = ' '
const AsciiEnd = '~'

func FileLines(symbols []byte) []string {
	return strings.Split(string(symbols), "\n")
}

func StoreInDictionary(alphabet []byte, line []string) map[byte][]string {
	q := 1
	m := make(map[byte][]string)          //map is used to store in a dictionary;
	for _, dictionary := range alphabet { //keys: single chars('A','c',etc.); values: a list with 8 items
		m[dictionary] = line[q : q+FontHeight] //One string for each of the 8 lines making up each ASCII symbol
		q = q + FontHeight + 1                 //1 + 8 + 1 + 8 ...
	}
	return m
}

func PrintSymbols(arg string, m map[byte][]string) {
	normalChars := strings.Split(arg, "\\n")
	res := ""
	for t := 0; t < len(normalChars); t++ {
		for i := 0; i < FontHeight; i++ {
			if normalChars[t] == "" {
				fmt.Println()
				break
			}
			for _, s := range normalChars[t] {
				if s < AsciiStart || s > AsciiEnd {
					s = '?' //if input string contains any non-ascii char, it prints '?'
				}
				res += m[byte(s)][i]
			}
			fmt.Println(res)
			res = ""
		}
	}
}
