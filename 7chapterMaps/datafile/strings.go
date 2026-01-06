package datafile

import (
	"bufio"
	"log"
	"os"
)

func GetStrings(fileName string) ([]string, error) {
	lines := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) // this obj is required for reading
	index := 0
	for scanner.Scan() {
		data := scanner.Text()
		lines = append(lines, data)
		index++
		// error can be here, but we don't read it here
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil { // returns empty arr and Err()
		return nil, scanner.Err()
	}
	return lines, nil

}
