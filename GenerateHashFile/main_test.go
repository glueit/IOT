package main

import (
	"encoding/hex"
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
