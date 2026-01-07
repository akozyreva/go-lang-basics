// Task is to count number of ocurrencies in file.
// 1 solution - with map (light solution)
// 2 solution - with arrays
package main

import (
	"7chapterMap/datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	// count the number of votes
	var names []string
	var counts []int
	countedVotes := make(map[string]int)
	for _, name := range lines {
		countedVotes[name]++
	}
	for _, name := range lines {
		isNameFound := false
		for _, sortedName := range names {
			if name == sortedName {
				isNameFound = true
				break
			}
		}
		if !isNameFound {
			names = append(names, name)
		}
	}
	for _, sortedName := range names {
		count := 0
		for _, name := range lines {
			if name == sortedName {
				count++
			}
		}
		counts = append(counts, count)
	}
	fmt.Println(names)
	fmt.Println(counts)
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
	fmt.Println(countedVotes)
	// for for map
	for name, countOfVotes := range countedVotes {
		fmt.Printf("%s: %d\n", name, countOfVotes)
	}
}
