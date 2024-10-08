package main

import (
	"bufio"
	"fmt"
	"os"
)

func printStatement (sentence string) string {
	return sentence
}

func main () {

	fmt.Println("Enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ :=  reader.ReadString('\n')
	fmt.Println(printStatement(sentence))

}