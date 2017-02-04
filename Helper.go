package main

import (
	"crypto/md5"
	"encoding/hex"
)

// ENCRYPT STRING WITH MD5
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
