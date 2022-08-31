package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

func WriteTxt(dir, filename, content string) {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("文件创建或打开失败", err)
	}
	filePath := path.Join(rootDir, dir)
	//err = os.MkdirAll(filePath+strconv.Itoa(time.Now().Year())+"/", os.ModePerm)
	err = os.MkdirAll(filePath+"/", os.ModePerm)
	if err != nil {
		fmt.Println("文件创建或打开失败", err)
	}
	fileAllPath := filePath + "/" + filename
	file, err := os.OpenFile(fileAllPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0)
	if err != nil {
		fmt.Println("文件创建或打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(content)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func ReadTxt(filepath string) string {
	result := ""

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	defer f.Close()

	li := bufio.NewReader(f)
	for {
		a, _, c := li.ReadLine()
		if c == io.EOF {
			break
		}
		result += string(a) + "\n"
	}
	return result
}
