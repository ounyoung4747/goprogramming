package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "cs r?cks~"
	replaceWords := strings.NewReplacer("?", "o")
	fixed := replaceWords.Replace(brokenWords)
	fmt.Println(fixed)

}
