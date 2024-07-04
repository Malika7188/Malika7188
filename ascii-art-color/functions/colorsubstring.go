package functions

import (
	"fmt"
	"strings"
)

func SubstringColoring(str []string, substr, color string) []string {

	colorCode := fmt.Sprintf("\033[%sm", color)
	resetCode := "\033[0m"
	colorstr := make([]string, len(str))

	//find the starting index of the substring
	for i, s := range str {
		startIndex := strings.Index(s, substr)
		if startIndex == -1 { //checks if the substring is available and if not it returns the original string
			colorstr[i] = s
			continue
		}
		initially := s[:startIndex]
		after := s[startIndex+len(substr):]
		colorstr[i] = initially + colorCode + substr + resetCode + after
		//coloredStr[i] = initially + colorCode + substr + resetCode + after
	}
	return colorstr
	//split the string into 3 parts: rhe innitial string, substring and the after string, which is te result after coloring

	//s:=[]string{"hey","jkt"}
}
