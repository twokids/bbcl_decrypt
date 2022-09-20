package util

import (
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type DynamicExcelBuilderDto struct {
	FileName string
	Columns  []string
	Values   [][]string
	Size     int
	BatchNo  string
}

//获取excel的行
func GetExcelRows(localPath string) ([][]string, error) {
	if !(strings.Contains(localPath, ".xlsx") ) {
		log.Fatalln("getFileRows 文件后缀异常")
		return nil, nil
	}
	f, err := excelize.OpenFile(localPath)
	if err != nil {
		log.Fatalln("Filetodblogic getExcelRows OpenFile 异常", err)
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Filetodblogic getExcelRows Close 异常", err)
		}
	}()
	sheetName:=f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalln("Filetodblogic getExcelRows GetRows 异常", err)
		return nil, err
	}
	return rows, nil
}

func GenerateExcel(dynamicInfo DynamicExcelBuilderDto) {
	rootDir, _ := os.Getwd()
	rootDir += "/docs/response/"

	//存储2022/22112221221.xlsx格式
	fileName := dynamicInfo.FileName + strconv.Itoa(time.Now().Year()) + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx"
	err := os.MkdirAll(rootDir+"/", os.ModePerm)
	if err != nil {
		return
	}
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		return
	}

	// Test max characters in a cell.
	sheethead := []interface{}{
	}
	for _, column := range dynamicInfo.Columns {
		sheethead = append(sheethead, column)
	}

	cell, _ := excelize.CoordinatesToCellName(1, 1)
	streamWriter.SetRow(cell, sheethead)

	style2, _ := file.NewStyle(&excelize.Style{
		Border: nil,
		Fill: excelize.Fill{
			Type:    "pattern", //纯色填充
			Pattern: 1,
			Color:   []string{"DFEBF6"},
			Shading: 0,
		},
		Font: nil,
		Alignment: &excelize.Alignment{
			Horizontal:      "center", //水平居中
			Indent:          0,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "", //垂直居中
			WrapText:        false,
		},
		Protection:    nil,
		NumFmt:        0,
		DecimalPlaces: 0,
		CustomNumFmt:  nil,
		Lang:          "",
		NegRed:        false,
	})
	_ = file.SetCellStyle("Sheet1", "A1", "A1", style2)

	for i, obj := range dynamicInfo.Values {
		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		tmp := []interface{}{}
		for _, s := range obj {
			tmp = append(tmp, s)
		}
		streamWriter.SetRow(cell, tmp)
	}

	streamWriter.Flush()
	// Save spreadsheet by the given path.

	save_path := filepath.Join(rootDir, fileName)
	file.SaveAs(save_path)
	return
}
