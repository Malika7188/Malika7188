package main

import (
	"fmt"
	"os"
	"strings"

	goreloaded "goreloaded/functions"
)

func main() {
	// call the argument from the commandline passed to the program and stores the arguments in arg variabe
	if len(os.Args) != 3 { // checks if number of commandline are not equal to 2 then prints error message
		fmt.Println("Usage: go run main.go <saple.txt> <result.txt>")
		os.Exit(1)
	}
	filename := os.Args[1]
	arg2 := os.Args[2]
	file, Error := os.ReadFile(filename) // read content of filename and return file containing the contents and the errors encountered
	if Error != nil {                    // prints the encountered errors then return to the main function
		fmt.Println(Error)
		return
	}
	// fmt.Println(string(file))
	words := strings.Split(string(file), " ") // split the content input string into word based on spaces
	for i := 0; i < len(words); i++ {
		if words[i] == "a" && goreloaded.CheckVowels(words[i+1][0]) {
			words[i] = "an"
		}
		if words[i] == "A" && goreloaded.CheckVowels((words[i+1][0])) {
			words[i] = "An"
		}
	}

	word1 := goreloaded.Hex(words)
	word2 := goreloaded.Bin(word1)
	word3 := goreloaded.Up(word2)
	word4 := goreloaded.Low(word3)
	word5 := goreloaded.Cap(word4)
	word6 := goreloaded.Punc(word5)

	word7 := goreloaded.HandleApostrophes(word6)

	// writing the files

	cue := (strings.Join(word7, " "))
	new := []byte(cue)
	// fmt.Println(words)
	os.WriteFile(arg2, new, 0o644)
}
