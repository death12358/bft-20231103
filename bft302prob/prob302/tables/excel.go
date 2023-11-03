package tables

// func (E ExcelData) GetDeadTable_xlsx() {
// 	// 打开Excel文件
// 	f, err := excelize.OpenFile("book1.xlsx")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 创建一个Map来存储数据
// 	dataMap := make(map[string][]string)

// 	// 获取第一个Sheet的名称
// 	sheetName := f.GetSheetName(0)

// 	// 获取第一个Sheet的所有行
// 	rows, err := f.GetRows(sheetName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 遍历每一行
// 	for _, row := range rows {
// 		// 将第一列的值作为Key
// 		key := row[0]
// 		dataMap[key] = make([]string, 0)

// 		// 将Key对应的后续列的数据存储到Map中
// 		for _, value := range row[1:] {
// 			dataMap[key] = append(dataMap[key], value)
// 		}
// 	}

// 	// 打印Map中的数据
// 	for key, value := range dataMap {
// 		fmt.Printf("Key: %s, Value: %v\n", key, value)
// 	}
// }
