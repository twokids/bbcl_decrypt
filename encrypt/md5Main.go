package encrypt

import (
	"bbcl-decrypt/util"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Md5Decrypt() {
	//t1 := time.Now()
	//elapsed1 := time.Since(t1)

	//读取excel文件内容
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

		_, fileName := filepath.Split(file)
		fileType := path.Ext(fileName)                         //文件类型
		fileNameOnly := strings.TrimSuffix(fileName, fileType) //文件名称，不带后缀
		dynamicInfo := util.DynamicExcelBuilderDto{
			FileName: fileNameOnly,
			Columns:  []string{"手机号MD5"},
			Values:   nil,
		}

		//读取excel的row。转化解析后的值
		onceCount := 500
		println("md5加密中，数据" + strconv.Itoa(len(rows)) + "~~~")
		for i := 0; i <= len(rows)/onceCount; i++ {
			t1 := time.Now()
			elapsed1 := time.Since(t1)

			//方法二
			curDynamicInfoChan := make(chan []string, 1000)
			md5DecryptSave(rows, i, onceCount, curDynamicInfoChan)
			close(curDynamicInfoChan) //释放ch
			for curV := range curDynamicInfoChan {
				dynamicInfo.Values = append(dynamicInfo.Values, curV)
			}

			elapsed1 = time.Since(t1)
			fmt.Printf("当前md5加密行参 onceCount:%v , i:%v , 执行时长:%v  \n", onceCount, i, elapsed1)
		}
		//保存到excel
		util.GenerateExcel(dynamicInfo)
	}
	return
}

func md5DecryptSave(rows [][]string, i int, onceCount int, curDynamicInfoChan chan []string) {
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 10) // 控制协程数量

	maxJCount := (i + 1) * onceCount
	if len(rows) > i*onceCount && len(rows) < (i+1)*onceCount {
		maxJCount = len(rows)
	}
	for j := i * onceCount; j < maxJCount; j++ {
		if i == 0 && j == 0 {
			//标题不处理
			continue
		}
		wg.Add(1)
		ch <- struct{}{}
		go func(row []string) {
			defer func() {
				<-ch
				wg.Done()
			}()
			curDynamicInfoChan <- util.Md5DecryptArray(row)
		}(rows[j])
	}
	wg.Wait()
	close(ch) //释放ch
	return
}
