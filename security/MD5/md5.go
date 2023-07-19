package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func main() {
	//hash := md5.New()
	hash := hmac.New(md5.New, []byte("secret"))
	bytes := []byte("hello\n")
	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()
	fmt.Printf("%x\n", hashValue)
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()

	hashValue256 := sha256.Sum256(bytes)
	fmt.Printf("%x", hashValue256)
}
