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
	fmt.Printf("%6s\n", "  请选择要执行的操作。")
	fmt.Printf("%6s\n", "  1，解密指定文件内容 。")
	fmt.Printf("%6s\n", "  21，加密指定文件内容 ")
	fmt.Printf("%6s\n", "  99，生成公钥私钥。")
	fmt.Printf("%6s\n", "***************************************")
	fmt.Println("请输入内容:")
}
func (c Welcome) RsaGenerateStep() {
	fmt.Printf("%6s\n", "***************************************")
	fmt.Printf("%6s\n", "   文件已生成至etc目录下")
	fmt.Printf("%6s\n", "   请妥善保管公钥和私钥")
	fmt.Printf("%6s\n", "   公钥提供给使用方")
	fmt.Printf("%6s\n", "   私钥用于解密密文")
	fmt.Printf("%6s\n", "***************************************")
	time.Sleep(time.Second * 6)
}

func (c Welcome) RsaEncryptStep() {
	fmt.Printf("%6s\n", "************************************************")
	fmt.Printf("%6s\n", "   请把需要加密的文件防至到docs/request文件夹下。 ")
	fmt.Printf("%6s\n", "   请确认密钥文件放至在etc文件夹下。")
	fmt.Printf("%6s\n", "   请等待，加密文件生成。")
	fmt.Printf("%6s\n", "   加密文件生成至docs/response文件夹下。")
	fmt.Printf("%6s\n", "************************************************")
	time.Sleep(time.Second * 8)
}


func (c Welcome) RsaDecryptStep() {
	fmt.Printf("%6s\n", "*************************************************")
	fmt.Printf("%6s\n", "   请把需要解密的文件防至到docs/request文件夹下。")
	fmt.Printf("%6s\n", "   请确认密钥文件放至在etc文件夹下。")
	fmt.Printf("%6s\n", "   请等待，解密文件生成。")
	fmt.Printf("%6s\n", "   加密文件生成至docs/response文件夹下。")
	fmt.Printf("%6s\n", "*************************************************")
	time.Sleep(time.Second * 8)
}
