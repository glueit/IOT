package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello world")
	file, err := os.Open("E:\\Personal\\Movies\\Theri (2016) HD DVDRip Tamil Full Movie Watch Online - www.TamilYogi.cc - Copy.mp4")

	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	outFile, err := os.Create("E:\\Personal\\Movies\\hash.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer outFile.Close()

	var offset int64
	chunkSize := 262144

	totalChunks := 0
	var sEnc []byte
	for {
		byteRead := make([]byte, chunkSize)
		totalBytesRead, err := file.ReadAt(byteRead, offset)
		if err != nil && err != io.EOF {
			fmt.Print(err)
			break
		}
		if err == io.EOF {
			fmt.Print("EOF reached")
			fmt.Printf("Last offset read %d", totalBytesRead)
			break
		}
		s256 := sha1.Sum(byteRead)
		s256bytes := s256[:]

		//sEnc := base64.URLEncoding.EncodeToString(by[:])
		//sEnc := by
		sEnc = append(sEnc, s256bytes...)
		//fmt.Fprintf(outFile, "%c", sEnc)
		totalChunks++
		offset += int64(totalBytesRead)
	}

	fmt.Println("Total Chunks", totalChunks)

	// filedata, _ := ioutil.ReadFile("E:\\Personal\\Movies\\sample.txt")
	// buffer := sha1.Sum(filedata)
	ioutil.WriteFile("E:\\Personal\\Movies\\hash.txt", sEnc, 0777)
	//fmt.Print(sEnc)

}
