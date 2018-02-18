package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	testStringBytes := []byte("Test String")
	expectedhash := "A5103F9C0B7D5FF69DDC38607C74E53D4AC120F2"
	hash := GenerateHash(testStringBytes)
	if hash == nil {
		t.Errorf("Hash obtained is null")
	}
	actualHash := strings.ToUpper(hex.EncodeToString(hash))
	result := strings.Compare(expectedhash, actualHash)
	if result != 0 {
		t.Errorf("Expected %s but found %s", expectedhash, actualHash)
	}
}

func TestGenerateHash256KB(t *testing.T) {
	testStringBytes := []byte("Test String")
	expectedhash := "A5103F9C0B7D5FF69DDC38607C74E53D4AC120F2"
	hash := GenerateHash(testStringBytes)
	if hash == nil {
		t.Errorf("Hash obtained is null")
	}
	actualHash := strings.ToUpper(hex.EncodeToString(hash))
	result := strings.Compare(expectedhash, actualHash)
	if result != 0 {
		t.Errorf("Expected %s but found %s", expectedhash, actualHash)
	}
}

func TestSliceTheDataUnEvenData(t *testing.T) {
	_ = ValidateDataSlicing(RandString(23), 2)
}

func TestSliceTheDataLessThanSliseSize(t *testing.T) {
	_ = ValidateDataSlicing(RandString(23), 31)
}

func TestSliceValidateZeroData(t *testing.T) {
	err := ValidateDataSlicing("", 2)
	if err == nil {
		t.Errorf("Expected error but did not find")
		return
	}
}

func TestSliceValidateZeroSliceLen(t *testing.T) {
	err := ValidateDataSlicing(RandString(20), 0)
	if err == nil {
		t.Errorf("Expected error but did not find")
		return
	}
}

func TestSliceValidateNegativeSliceLen(t *testing.T) {
	err := ValidateDataSlicing(RandString(20), -1)
	if err == nil {
		t.Errorf("Expected error but did not find")
		return
	}
}

func ValidateDataSlicing(testString string, slicingSize int) error {

	testStringBytes := []byte(testString)
	slicedData, err := SliceTheData(testStringBytes, slicingSize)
	actualDataSize := len(slicedData)

	if err != nil {
		return err
	}

	expectedDataSize := GetExpectedDataSize(testStringBytes, slicingSize)
	if expectedDataSize != actualDataSize {
		errorMessage := fmt.Sprintf("Expected data size %d but found %d", expectedDataSize, actualDataSize)
		return errors.New(errorMessage)
	}
	return nil

}

func GetExpectedDataSize(data []byte, slicesize int) int {
	var expectedDataSize int
	dataLength := len(data)
	remainder := dataLength % slicesize

	expectedDataSize = dataLength / slicesize

	if remainder != 0 {
		expectedDataSize++
	}
	return expectedDataSize
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
