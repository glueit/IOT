package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

var inFile = flag.String("Infile", "E:\\Personal\\Movies\\Theri (2016) HD DVDRip Tamil Full Movie Watch Online - www.TamilYogi.cc - Copy.mp4", "File Name to hash")
var outFolder = flag.String("OutFolder", ".", "Output folder")

func main() {
	rawData, err := ReadData()
	check(err)
	slicedData, err := SliceTheData(rawData, 262144)
	check(err)

	fileName := GetFileName(rawData)
	fullpath := path.Join(*outFolder, fileName)

	generatedHashes := GenerateHash(slicedData)
	ioutil.WriteFile(fullpath, generatedHashes, 0777)
	fmt.Println("hash file generated successfull - ", fileName)
}

func GetFileName(rawData []byte) string {
	var data [][]byte
	data = append(data, rawData)
	var fileHash = GenerateHash(data)
	return strings.ToUpper(hex.EncodeToString(fileHash)) + ".hash"
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
	data, err := ioutil.ReadFile(fmt.Sprintf("%s", *inFile))
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
