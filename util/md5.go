package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 返回一个32位md5加密后的字符串
func Md5Decrypt(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5DecryptArray(arr []string) []string {
	result := []string{}
	for _, str := range arr {
		h := md5.New()
		h.Write([]byte(str))
		result = append(result, hex.EncodeToString(h.Sum(nil)))
	}
	return result
}
