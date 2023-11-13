package main

import "github.com/xuri/excelize/v2"

func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("壁画知识点0523", "B2")
	if err != nil {
		println(err.Error())
		return
	}
	println("+++", cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("壁画知识点0523")
	for _, row := range rows {
		for _, colCell := range row {
			print(colCell, "\t===")
		}
		println("||")
	}
}
