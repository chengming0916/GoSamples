package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {

	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 读取文件流
	// file1, err := os.Open("")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file1.Close()
	// excelize.OpenReader(file1)

	// 创建工作表
	index, err := file.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置单元格
	file.SetCellValue("Sheet2", "A2", "Hello world")
	file.SetCellValue("Sheet1", "B2", 100)

	// 设置工作薄的默认工作表
	file.SetActiveSheet(index)

	_, err = file.NewSheet("Sheet3")
	if err != nil {
		fmt.Println(err)
		return
	}
	values := make([]int, 9)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			values[j-1] = i * j
		}

		// 按行赋值
		err = file.SetSheetRow("Sheet3", fmt.Sprintf("A%d", i), &values)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 按列赋值
	// err = file.SetSheetCol("Sheet3", "J1", &[]any{})

	// 读取指定单元格
	file.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}

	//读取Sheet1上所有单元格
	rows, err := file.Rows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
		}
		for _, cell := range row {
			fmt.Println(cell, "\t")
		}
		fmt.Println()
	}

	if err = rows.Close(); err != nil {
		fmt.Println(err)
	}
	// for _, row := range rows {
	// 	for _, cell := range row {
	// 		fmt.Println(cell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	if err := file.SaveAs("Book-1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
