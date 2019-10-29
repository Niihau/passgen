package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/an63/gut/yrand"
)

var (
	count          = flag.Uint("c", 10, "number of password candidates")
	length         = flag.Uint("l", 16, "length of each password")
	hasDigit       = flag.Bool("1", true, "whether to include numerical digits")
	hasLowerLetter = flag.Bool("a", true, "whether to include letters in lower case")
	hasUpperLetter = flag.Bool("A", true, "whether to include letters in upper case")
	specialChars   = flag.String("s", "-_", "special chars to include")
)

func main() {
	flag.Parse()

	sb := strings.Builder{}
	if *hasLowerLetter {
		sb.WriteString("abcdefghijkmnopqrstuvwxyz")
	}
	if *hasUpperLetter {
		sb.WriteString("ABCDEFGHJKLMNPQRSTUVWXYZ")
	}
	if *hasDigit {
		sb.WriteString("23456789")
	}
	if len(*specialChars) > 0 {
		sb.WriteString(*specialChars)
	}

	allChars := sb.String()
	pwdCount := int(*count)
	pwdLength := int(*length)

	if len(allChars) <= 0 {
		fmt.Println("got no candidate chars to generate password")
		os.Exit(1)
	} else if pwdCount <= 0 {
		fmt.Println("count should be at least 1")
		os.Exit(2)
	} else if pwdLength <= 0 {
		fmt.Println("length should be at least 1")
		os.Exit(3)
	}

	for i := 0; i < pwdCount; i++ {
		key, err := yrand.Runes(allChars, pwdLength)
		if err != nil {
			fmt.Println("got error:", err)
			return
		}

		fmt.Println(key)
	}
}
