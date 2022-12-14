package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"testing"
)

func TestName(t *testing.T) {

	//公钥加密
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("文件创建或打开失败", err)
	}
	//filePath := path.Join(rootDir, "../etc/public.txt")
	inputPublicStr := ReadTxt("../etc/public.txt")
	publicKey, _ := ParseRsaPublicKeyFromPemStr(inputPublicStr)
	s:="nanshan南山！！~~%￥%……你好山东省聊城市莘县莘县徒骇河南岸通运0933secret message11莘县莘dddddddddddddddddddddd县徒"
	if len(s) > 124 {
		s = s[len(s)-124:]
	}
	encryptedBytes, err := rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		publicKey,
		[]byte(s),
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("encrypted bytes: ", encryptedBytes)
	str1 := base64.RawURLEncoding.EncodeToString(encryptedBytes) + "tx2022"
	fmt.Println("encrypted str1: ", str1)

	str2, _ := base64.RawURLEncoding.DecodeString(strings.Split(str1, "tx2022")[0])
	encryptedBytesAgain := []byte(str2)
	fmt.Println("encrypted encryptedBytesAgain: ", encryptedBytesAgain)

	//私钥解密
	filePath2 := path.Join(rootDir, "./etc/private.txt")
	inputPrivateStr := ReadTxt(filePath2)
	privateKey, _ := ParseRsaPrivateKeyFromPemStr(inputPrivateStr)
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytesAgain, &rsa.OAEPOptions{Hash: crypto.SHA512})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("decrypted message: ", string(decryptedBytes))
}
