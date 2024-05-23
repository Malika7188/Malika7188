package revision

import (
	"fmt"
)

func Chunk(slice []int, size int) {
	mySlice := [][]int{}

	if size == 0 {
		return
	}
	for fastIndex := 0; fastIndex < len(slice)-1; fastIndex += size {
		lastIndex := fastIndex + size
		if lastIndex > len(slice) {
			lastIndex = len(slice)
		}
		mySlice = append(mySlice, slice[fastIndex:lastIndex])
	}
	fmt.Println(mySlice)
}
