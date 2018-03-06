package main

import (
	"fmt"
	"log"
	"testing"
)

func TestReadYamlFile(t *testing.T) {
	yamlFile, err := readYamlFile()

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	} else {
		fmt.Println("File read successful")
	}

	fmt.Println(yamlFile)
}
