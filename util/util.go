package util

import (
	"fmt"
	"hash/crc32"
	"log"
)

// Hash a string using CRC32
func Hash(link string) string {
	num := crc32.ChecksumIEEE([]byte(link))
	return fmt.Sprintf("%x", num)
}

// CheckError is not nil, and ends the program
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
