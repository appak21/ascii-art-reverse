package main

import (
	ascii "ascii/pkg"
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	alphabet, err := ioutil.ReadFile("pkg/ascii-chars.txt")
	Check(err)
	symbols := make([]string, 0, 855)
	artFile, err := os.Open("pkg/banner.txt")
	Check(err)
	defer artFile.Close()
	fileScanner := bufio.NewScanner(artFile)
	for fileScanner.Scan() {
		symbols = append(symbols, fileScanner.Text())
	}
	dictionary := ascii.StoreInDictionary(alphabet, symbols)
	fileName := flag.String("reverse", "", "Use flag --reverse=<fileName>")
	flag.Parse()
	if len(flag.Args()) == 1 {
		ascii.PrintSymbols(flag.Arg(0), dictionary)
	}
	if *fileName == "" {
		return
	}
	artReverse(fileName, dictionary)
}

func artReverse(fileName *string, dictionary map[byte][]string) {
	givenSymbols, err := ioutil.ReadFile(*fileName)
	Check(err)
	givenLine := ascii.FileLines(givenSymbols)

	res := ""
	for k := 0; k < len(givenLine)/8; k++ {
		symbKey := reversedKey(k*8, givenLine)
		var arr []byte
		for i := 0; i < len(symbKey); i++ {
			for i2, v2 := range dictionary {
				if strings.Join(symbKey[byte(i)], "\n") == strings.Join(v2, "\n") {
					arr = append(arr, i2)
				}
			}
		}
		if len(arr) != len(symbKey) || len(arr) == 0 {
			fmt.Println("Your banner format isn't correct")
			return
		}
		res += string(arr)
		if k != len(givenLine)-1 {
			res += "\n"
		}
	}
	fmt.Print(res)
}

func Check(e error) {
	if e != nil {
		log.Fatal("couldn't open the file")
	}
}

func maxLen(el int, line []string) (int, []string) {
	maxlen := 0
	minlen := len(line[el])
	t := 0
	for i := el; i < el+ascii.FontHeight; i++ {
		l := len(line[i])
		if l > maxlen {
			maxlen = l //got the max length of the lines.
			t = i
		}
		if i < 6 && l < minlen {
			minlen = l
		}
	}
	if maxlen >= 5 && line[t][maxlen-5:maxlen] == "     " && minlen != maxlen { //if all 8 lines don't contain the same space
		return 0, nil
	}
	maxlen = maxlen + 1
	for i := el; i < el+ascii.FontHeight; i++ {
		for len(line[i]) < maxlen {
			line[i] += " " //got matched with the standard format
		}
	} //now we got our lines corrected as we needed
	return maxlen, line
}

func reversedKey(el int, line []string) map[byte][]string { //this function receives 8 lines and returns given symbols as keys and values
	symbKey := make(map[byte][]string) //the number of keys is equal to the number of symbols in the given file
	key := 0                           // 0 key contains the first symbol, 1 key - the second, so on...
	maxlen, line := maxLen(el, line)
	count := 0
	m := 0
	for j := 0; j < maxlen; j++ {
		for i := el; i < el+ascii.FontHeight; i++ {
			if line[i][j] == ' ' { //symbols were taken separately by a free space that takes one column
				count++
			} else {
				count = 0
				break
			}
			if count == 8 {
				count = 0
				if m+6 <= maxlen && lookForSpaces(m, el, line) { //a space symbol consists of 6 normal spaces. In order to check if there's a space symbol, we need to be sure that next 6 normal spaces exist
					j = j + 5 //jumps 6 spaces to avoid checking next column
				} else if m == j {
					m = j + 1
					break
				}
				symbKey[byte(key)] = applygivenLine(m, j+1, el, line) //m is the start, j+1 is the end of each lines, which make one symbol
				m = j + 1
				key++
			}
		}
	}
	return symbKey
}

func applygivenLine(m, n, el int, line []string) []string {
	givenLine := make([]string, ascii.FontHeight)
	for i := 0; i < ascii.FontHeight; i++ {
		givenLine[i] = line[i+el][m:n]
	}
	return givenLine
}

func lookForSpaces(m, el int, line []string) bool {
	for i := el; i < el+ascii.FontHeight; i++ {
		if line[i][m:m+6] != "      " {
			return false
		}
	}
	return true
}
