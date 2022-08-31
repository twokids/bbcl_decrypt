package main

import (
	"bbcl-decrypt/util"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	var wel util.Welcome
	wel.Welcome()

	wel.OpeartionStep()
	step1 := util.GetStrInput()
	if step1 == "1" {
		wel.RsaDecryptStep()
		rsaDecrypt()
	} else if step1 == "2" {
		wel.RsaGenerateStep()
		rsaGenerate()
	} else {
		fmt.Printf("%6s\n", "***  信息输入错误。 ***")
		wel.Welcome()
	}
}

func rsaDecrypt() {
	//获取私钥
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("文件创建或打开失败", err)
	}
	filePath1 := path.Join(rootDir, "./etc/public.txt")
	inputPublicStr := util.ReadTxt(filePath1)
	publicKey, _ := util.ParseRsaPublicKeyFromPemStr(inputPublicStr)

	filePath := path.Join(rootDir, "./etc/private.txt")
	inputPrivateStr := util.ReadTxt(filePath)
	privateKey, _ := util.ParseRsaPrivateKeyFromPemStr(inputPrivateStr)
	println(publicKey, privateKey)

	//扫描指定文件夹
	dirPath, _ := filepath.Abs("./docs/request")

	//1,扫描需要处理的request下的内容
	//读取dirpath路径下文件信息，循环处理
	files, _ := filepath.Glob(dirPath + "/*")
	for _, file := range files {
		if file == dirPath {
			continue
		}
		//单个文件处理
		rows, err := util.GetExcelRows(file)
		if err != nil {
			log.Fatalln("getExcelRows获取文件内容异常", err)
			return
		}

		if rows == nil {
			continue
		}

		//拼接head
		var headCols []string
		for _, v1 := range rows[0] {
			headCols = append(headCols, v1)
		}

		_, fileName := filepath.Split(file)
		fileType := path.Ext(fileName)                         //文件类型
		fileNameOnly := strings.TrimSuffix(fileName, fileType) //文件名称，不带后缀
		dynamicInfo := util.DynamicExcelBuilderDto{
			FileName: fileNameOnly,
			Columns:  headCols,
			Values:   nil,
		}

		//读取excel的row。转化解析后的值
		onceCount := 500
		for i := 0; i <= len(rows)/onceCount; i++ {
			maxJCount := (i + 1) * onceCount
			if len(rows) > i*onceCount && len(rows) < (i+1)*onceCount {
				maxJCount = len(rows)
			}
			for j := i * onceCount; j < maxJCount; j++ {
				if i == 0 && j == 0 {
					//标题不处理
					continue
				}
				dynamicInfo.Values = append(dynamicInfo.Values, util.Decrypt(privateKey, rows[j]))
				//dynamicInfo.Values = append(dynamicInfo.Values, util.Encrypt(publicKey, rows[j]))
			}
		}
		//保存到excel
		util.GenerateExcel(dynamicInfo)
	}
	return
}

//生成密钥文件
func rsaGenerate() {
	// Create the keys
	priv, pub := util.GenerateRsaKeyPair()
	// Export the keys to pem string
	priv_pem := util.ExportRsaPrivateKeyAsPemStr(priv)
	pub_pem, _ := util.ExportRsaPublicKeyAsPemStr(pub)
	util.WriteTxt("./etc", "private.txt", priv_pem)
	util.WriteTxt("./etc", "public.txt", pub_pem)
}

//加解密示例demo
func rsaEncryptDemo() {
	//公钥加密
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("文件创建或打开失败", err)
	}
	filePath := path.Join(rootDir, "./etc/public.txt")
	inputPublicStr := util.ReadTxt(filePath)
	publicKey, _ := util.ParseRsaPublicKeyFromPemStr(inputPublicStr)

	encryptedBytes, err := rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		publicKey,
		[]byte("nanshan南山！！~~%￥%……你好0933secret message"),
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("encrypted bytes: ", encryptedBytes)
	str1 := base64.RawURLEncoding.EncodeToString(encryptedBytes) + "tx2022"
	fmt.Println("encrypted str1: ", str1)

	str2,_:=base64.RawURLEncoding.DecodeString(strings.Split(str1, "tx2022")[0])
	encryptedBytesAgain := []byte(str2)
	fmt.Println("encrypted encryptedBytesAgain: ", encryptedBytesAgain)

	//私钥解密
	filePath2 := path.Join(rootDir, "./etc/private.txt")
	inputPrivateStr := util.ReadTxt(filePath2)
	privateKey, _ := util.ParseRsaPrivateKeyFromPemStr(inputPrivateStr)
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytesAgain, &rsa.OAEPOptions{Hash: crypto.SHA512})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("decrypted message: ", string(decryptedBytes))
}
