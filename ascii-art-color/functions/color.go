package functions

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ColoredSub(substr string, fontarray [95][8]string, colorcode string) []string {
	concatenatedLines := make([]string, 8)

	for _, char := range substr {
		Ascii, err := PrintChar(char, fontarray)
		if err != nil {
			log.Fatal(err)
		}
		for i, line := range Ascii {
			if colorcode != "" {
				concatenatedLines[i] += fmt.Sprintf("\033[%sm%s\033[0m", colorcode, line)
			} else {
				concatenatedLines[i] += line
			}
		}
	}
	return concatenatedLines
}

func RgbToAnsi(color string) string {
	color = strings.ToLower(color)

	colr := strings.TrimPrefix(color, "(")
	colr1 := strings.TrimPrefix(colr, ")")
	colr2 := strings.Split(colr1, ",")

	red, err := strconv.Atoi(colr2[0])
	CheckError(err)
	green, err := strconv.Atoi(colr2[1])
	CheckError(err)
	blue, err := strconv.Atoi(colr2[2])
	CheckError(err)
	Colors := fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue)

	return Colors
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
