package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	mathskils "mathskils/functions"
)

func main() {
	if len(os.Args) != 2 { // reading the command-lne arguments
		fmt.Println("usage: go run <main.go><data.txt>")
		return
	}

	filedata := os.Args[1]
	file, err := os.Open(filedata) // opening the argument at index 1 for scanning
	if err != nil {                // checking the occurence of errors
		fmt.Println(err)
		return
	}
	var values []float64 // initializing variable to store the final data read from filedat

	scanner := bufio.NewScanner(file) // scan the filedata line by line
	for scanner.Scan() {
		filecontent := scanner.Text()
		if scanner.Text() == "" { // checks if the file data is empty
			return
		}
		num, err := strconv.ParseFloat(filecontent, 64) // converts the filedata contents to float
		if err != nil {
			fmt.Println(err)
			return
		}
		values = append(values, num) // append the data read from filedata in value variable
	}
	// calling all the math skills functions
	output := mathskils.HandleAverage(values)
	output1 := mathskils.HandleMedian(values)
	output2 := mathskils.HandleStandardDeviation(values)
	output3 := mathskils.HandleVariance(values)
	// formating all the math funtions
	fmt.Printf("Average: %.0f\n", output)
	fmt.Printf("Medium: %.0f\n", output1)
	fmt.Printf("Standard Deviation: %.0f\n", output2)
	fmt.Printf("Variance: %.0f\n", output3)
}
