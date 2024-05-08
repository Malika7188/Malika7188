package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	mathskils "mathskils/functions"
)

func main() {
	// reading the number of arguments from the comandline
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . data.txt")
		os.Exit(1)
	}
	arg := os.Args[1]
	file, err := os.ReadFile(arg)
	if err != nil {
		fmt.Println("error encountered")
		return
	}
	content := strings.Split(string(file), "\n")
	var values []float64
	for _, char := range content {
		if char == "" {
			continue
		}
		// parse each character as a float
		num, err := strconv.ParseFloat(char, 64)
		if err != nil {
			fmt.Println("error encountered")
			os.Exit(1)
		}
		values = append(values, num)
	}

	output := mathskils.HandleAverage(values)
	output1 := mathskils.HandleMedian(values)
	output2 := mathskils.HandleVariance(values)
	output3 := mathskils.HandleStandardDeviation(values)

	fmt.Println("Average:", math.Round(output))
	fmt.Println("Median:", math.Round(output1))
	fmt.Println("Variance:", math.Round(output2))
	fmt.Println("Standard Deviatiation:", math.Round(output3))
}
