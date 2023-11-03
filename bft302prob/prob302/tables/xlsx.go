package tables

import (
	"log"

	"github.com/xuri/excelize/v2"
)

// const (
// 	dataKeyIndex = 0
// 	gameDir      = "gametemplate"
// 	dataDir      = "datas"
// )

// func init() {
// 	currentDir, _ := os.Getwd() //當下目錄
// 	// LogTool.LogInfo("slot init", "current dir:", currentDir)
// 	currentDir = strings.ReplaceAll(currentDir, "\\", "/")
// 	//計算返回目標資料夾需要幾階".."
// 	dirSlice := strings.Split(currentDir, "/")
// 	numDotDot := 0
// 	for l := len(dirSlice) - 1; l > 0; l-- {
// 		if dirSlice[l] == gameDir {
// 			break
// 		}
// 		numDotDot++
// 	}
// 	elem := generateDotDotSlice(numDotDot)
// 	elem = append([]string{currentDir}, elem...)

// 	// 使用..来回到上一级目录，然后继续返回上级目录，直到回到gameDir
// 	parentDir := filepath.Join(elem...)

// 	// 產生進入datas的路徑
// 	targetDir := filepath.Join(parentDir, dataDir)

// 	//設定路徑
// 	mappingPath = strings.ReplaceAll(targetDir+"\\*\\config\\mapping.csv", "\\", "/")
// 	LogTool.LogInfo("slot init", "mappingPath:", mappingPath)

// 	payTablePath = strings.ReplaceAll(targetDir+"\\*\\pay\\table.csv", "\\", "/")
// 	LogTool.LogInfo("slot init", "payTablePath:", payTablePath)

// 	payLinesPath = strings.ReplaceAll(targetDir+"\\*\\pay\\lines.csv", "\\", "/")
// 	LogTool.LogInfo("slot init", "payLinesPath:", payLinesPath)

// 	probabilityPath = strings.ReplaceAll(targetDir+"\\*\\probability\\", "\\", "/")
// 	LogTool.LogInfo("slot init", "probabilityPath:", probabilityPath)
// }

// var (
// 	mappingPath     string
// 	payTablePath    string
// 	payLinesPath    string
// 	probabilityPath string
// )

// func generateDotDotSlice(n int) []string {
// 	if n <= 0 {
// 		return nil
// 	}
// 	s := make([]string, n)
// 	for i := range s {
// 		s[i] = ".."
// 	}
// 	return s
// }

type Folder string

func TableInit() {
	// DeadProb會用payTable中的值去算, 所以順序不能變
	GetFishPayTable()
	GetFishDeadProb()
}

var DeadMap DeadProbMap_flow

// Sheet --> 資料內容(map)
type ExcelData map[string]DataMap

// 數據名稱(first col) --> 數據內容([]string)
type DataMap map[string][]string

// 好像要調整成os.Executable()?
// fileName: "path/XXX.xlsx"
//
// ...............................	map: Sheet名稱 --> 檔案內容(map)
func GetExcelData(fileName string) (excelData ExcelData) {
	excelData = map[string]DataMap{}
	// 打开Excel文件
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// 獲取所有Sheet
	for _, sheet := range f.GetSheetList() {
		//創建map紀錄資料
		dataMap := make(map[string][]string)

		if _, ok := excelData[sheet]; !ok {
			excelData[sheet] = make(map[string][]string)
		}

		rows, err := f.GetRows(sheet)
		if err != nil {
			log.Fatal(err)
		}
		// 遍历每一行
		for _, row := range rows {
			// 将第一列的值作为Key
			key := row[0]
			dataMap[key] = make([]string, 0)
			// 将Key对应的后续列的数据存储到Map中
			for _, value := range row[1:] {
				dataMap[key] = append(dataMap[key], value)
			}
		}
		excelData[sheet] = dataMap
	}

	// // 打印Map中的数据

	// js, _ := json.Marshal(excelData)

	// fmt.Printf("PayTableMap:\n%#v", string(js))
	// fmt.Printf("End GetExcelData(fileName string:\n")
	return
}

// func main() {
// 	// fmt.Printf("\nGetExcelData(\"payTable.xlsx\")~~~~~\n%#v\n~~~~~\n", GetExcelData("payTable.xlsx"))
// 	// E := GetExcelData("deadTable.xlsx")
// 	// fmt.Printf("\nGetExcelData(\"deadTable.xlsx\")~~~~~\n%#v\n~~~~~\n", GetExcelData("deadTable.xlsx"))
// 	// js, err := json.Marshal(E.GetDeadTableMap())
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Printf("\nE.GetDeadTableMap()~~~~~\n%#v\n~~~~~\n", string(js))

// 	E2 := GetExcelData("payTable.xlsx")
// 	js2, err := json.Marshal(E2.GetPayTableMap())
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("\nE2.GetPayTableMap()~~~~~\n%#v\n~~~~~\n", string(js2))
// }

// ChartWeight             []int32 // 表權重
// SysWin_AdjustMultiplier map[string]float64
// TableWeight             map[string]([]int32)
// IntervalWeights         map[string]([]int32) // 區間權重
// PointInterval           map[string]([]int64)

//	func NewPayTable() (payTable *PayTable) {
//		payTable = &PayTable{
//			ChartWeight:             []int32{},
//			SysWin_AdjustMultiplier: map[string]float64{},
//			DataMap:                 map[string][]string{},
//			TableWeight:             map[string]([]int32){},
//			IntervalWeights:         map[string]([]int32){}, // 區間權重
//			PointInterval:           map[string]([]int64){},
//		}
//		return
//	}
