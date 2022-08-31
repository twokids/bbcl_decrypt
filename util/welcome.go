package util

import (
	"fmt"
	"time"
)

type Welcome struct {
}

func (c Welcome) Welcome() {
	fmt.Printf("%6s\n", "***************************************")
	fmt.Printf("%6s\n", "***        同犀解密小工具欢迎您          ***")
	fmt.Printf("%6s\n", "***************************************")
	time.Sleep(time.Second * 2)
}

func (c Welcome) OpeartionStep() {
	fmt.Printf("%6s\n", "***************************************")
	fmt.Printf("%6s\n", "***  请选择要执行的操作。   ***")
	fmt.Printf("%6s\n", "***  1，解密指定文件内容 。 ***")
	fmt.Printf("%6s\n", "***  2，生成公钥私钥。     ***")
	fmt.Printf("%6s\n", "***************************************")
	fmt.Println("请输入内容:")
}
func (c Welcome) RsaGenerateStep() {
	fmt.Printf("%6s\n", "***************************************")
	fmt.Printf("%6s\n", "***  文件已生成                       ***")
	fmt.Printf("%6s\n", "***  请妥善保管公钥和私钥                ***")
	fmt.Printf("%6s\n", "***  公钥提供给使用方                   ***")
	fmt.Printf("%6s\n", "***  私钥用于解密密文                   ***")
	fmt.Printf("%6s\n", "***************************************")
	time.Sleep(time.Second * 2)
}

func (c Welcome)RsaDecryptStep() {
	fmt.Printf("%6s\n", "***************************************************")
	fmt.Printf("%6s\n", "***  请把需要解密的文件防至到decrypt/request文件夹下。 ***")
	fmt.Printf("%6s\n", "***  请把密钥文件放至到decrypt文件夹下。              ***")
	fmt.Printf("%6s\n", "***  请等待，解密文件生成。                          ***")
	fmt.Printf("%6s\n", "***************************************************")
	time.Sleep(time.Second * 2)
}
