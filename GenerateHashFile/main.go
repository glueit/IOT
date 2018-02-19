package main

import (
	"crypto/sha1"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
)

var file = flag.String("file", "E:\\Personal\\Movies\\Theri (2016) HD DVDRip Tamil Full Movie Watch Online - www.TamilYogi.cc - Copy.mp4", "Port to connect to")

func main() {
	rawData, err := ReadData()
	check(err)
	slicedData, err := SliceTheData(rawData, 262144)
	check(err)
	generatedHashes := GenerateHash(slicedData)
	fmt.Print(generatedHashes)
}

func GenerateHash(slicedData [][]byte) []byte {
	var completeHash []byte
	for i := 0; i < len(slicedData); i++ {
		hash := sha1.Sum(slicedData[i])
		completeHash = append(completeHash, hash[:]...)
	}
	return completeHash
}

func ReadData() ([]byte, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s", file))
	if err != nil {
		fmt.Printf("Help message: <executable> -file <file path>")
		return nil, err
	}
	return data, nil
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
