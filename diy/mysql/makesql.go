package main

import (
	"fmt"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// 文件路径
const filePath = "/root/code-project/go-project/diy/mysql/output.xlsx"
// sheet名称
const sheetName = "Sheet"
// 输出文件路径
const outputPath = "output.txt"
// id列
const idCol = 1
// encrypt_attach列
const encryptAttach = 4
// sql
const sql = "update charge_channel set encrypted_attach = '%s' where id = %s;"

func main() {
	// Open the Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the values from the first and tenth columns
	rows := f.GetRows(sheetName)
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i, row := range rows {
		if i == 0 {
			continue
		}
		id := row[idCol - 1]
		encryptedAttach := row[encryptAttach - 1]

		// Print the SQL statement with the values
		sql := fmt.Sprintf(sql, encryptedAttach, id)
		fmt.Fprintln(file, sql)
	}
}
