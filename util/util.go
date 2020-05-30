package util

import (
	"fmt"
	"hash/crc32"
	"log"
)

// Link holds information about any link stored in the database
type Link struct {
	ID   string `json:"id"`
	Link string `json:"link"`
}

// Result holds information that could be returned from any request
type Result struct {
	Link  *Link   `json:"link,omitempty"`
	Error *string `json:"error,omitempty"`
}

// Hash a string using CRC32
func Hash(link string) string {
	num := crc32.ChecksumIEEE([]byte(link))
	return fmt.Sprintf("%x", num)
}

// CheckError is not nil, and ends the program
func CheckError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// StringPtr creates a pointer from a string
func StringPtr(str string) *string {
	return &str
}
