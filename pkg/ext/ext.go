package ext

import (
	"crypto/md5"
	"encoding/hex"
)

func InString(filed string, array []string) bool {
	for _, str := range array {
		if filed == str {
			return true
		}
	}
	return false
}

func Md5String(str string) string {
	hash := md5.Sum([]byte(str))
	md5String := hex.EncodeToString(hash[:])
	return md5String
}
