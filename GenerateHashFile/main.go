package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
)

func main() {
	fmt.Print("Hello")
}

func GenerateHash(data []byte) []byte {
	hash := sha1.Sum(data)
	return hash[:]
}

func SliceTheData(data []byte, sliceSize int) ([][]byte, error) {

	totalDataLength := len(data)
	var index int
	var slicedData [][]byte
	var exception error

	if sliceSize <= 0 {
		fmt.Print("sliceSize should be greater than 0")
		exception = errors.New("sliceSize should be greater than 0")
		return nil, exception
	}

	if totalDataLength == 0 {
		fmt.Print("Data size is 0")
		exception = errors.New("Data size is 0")
		return nil, exception
	}

	for totalDataLength >= (index + sliceSize) {
		slicedData = append(slicedData, data[index:(index+sliceSize)])
		index += sliceSize
	}

	if totalDataLength > index {
		slicedData = append(slicedData, data[index:totalDataLength-1])
	}

	return slicedData, nil
}
