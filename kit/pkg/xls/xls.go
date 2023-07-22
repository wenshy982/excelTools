package xls

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"

	"github.com/xuri/excelize/v2"

	"kit/tools/osx"
)

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

type XLS struct {
	Input     string // 输入文件
	Output    string // 输出文件
	Template  string // 模板文件
	Errors    string // 错误文件
	OutputDir string // 输出目录
	OutPutCSV string // 输出 CSV 文件
}

// New 新建 Excel 并设置可选项
func New(opts ...Opt) *XLS {
	var o = &XLS{}
	for _, opt := range DefaultOpt {
		opt(o)
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Copy 拷贝文件
func (e *XLS) Copy(src, dst string) {
	osx.CopyFile(src, dst)
	log.Printf("【拷贝文件】拷贝完成，源文件：%s，目标文件：%s \n", src, dst)
}

// CopyInputTo 复制输入文件到指定目录
func (e *XLS) CopyInputTo(dir string) {
	e.Copy(e.Input, dir)
	log.Printf("【复制输入文件】复制完成，输入文件：%s，输出文件：%s \n", e.Input, dir)
}

// CopyTemplateTo 复制模板文件到指定目录
func (e *XLS) CopyTemplateTo(dir string) {
	e.Copy(e.Template, dir)
	log.Printf("【复制模板文件】复制完成，模板文件：%s，输出文件：%s \n", e.Template, dir)
}

// CopyInputToOutput 复制输入文件到输出文件
func (e *XLS) CopyInputToOutput() {
	e.CopyInputTo(e.Output)
}

// CopyTemplateToOutput 复制模板文件到输出文件
func (e *XLS) CopyTemplateToOutput() {
	e.CopyTemplateTo(e.Output)
}

// NewRow 将字符串数组转换为接口数组
func (e *XLS) NewRow(values ...string) []interface{} {
	result := make([]interface{}, len(values))

	for i, value := range values {
		switch {
		case value == "":
			result[i] = nil
		case isInt(value):
			intValue, _ := strconv.Atoi(value)
			result[i] = intValue
		case isFloat(value):
			floatValue, _ := strconv.ParseFloat(value, 64)
			result[i] = floatValue
		default:
			result[i] = value
		}
	}

	return result
}

// WriteToRow 将接口数组写入到指定行
func (e *XLS) WriteToRow(f *excelize.File, sheet string, row []interface{}, oRowsLen int) {
	_ = f.SetSheetRow(sheet, "A"+strconv.Itoa(oRowsLen), &row)
}

// InitOutputDir 初始化输出目录
func (e *XLS) InitOutputDir() {
	osx.InitDir(e.OutputDir)
	log.Println("【初始化输出目录】初始化完成，输出目录：", e.OutputDir)
}

// Open 打开文件
func (e *XLS) Open(path string) (f *excelize.File) {
	f, _ = excelize.OpenFile(path)
	return
}

// OpenInput 打开输入文件
func (e *XLS) OpenInput() (f *excelize.File) {
	return e.Open(e.Input)
}

// OpenOutput 打开输出文件
func (e *XLS) OpenOutput() (f *excelize.File) {
	return e.Open(e.Output)
}

// OpenTemplate 打开模板文件
func (e *XLS) OpenTemplate() (f *excelize.File) {
	return e.Open(e.Template)
}

// OpenErrors 打开错误文件
func (e *XLS) OpenErrors() (f *excelize.File) {
	return e.Open(e.Errors)
}

// Save 保存文件
func (e *XLS) Save(f *excelize.File) {
	_ = f.Save()
}

// SaveAsCSV 保存为 CSV 文件
func (e *XLS) SaveAsCSV(input string, output string) {
	iFile := e.Open(input)
	defer func() { _ = iFile.Close() }()
	_, iRows := e.Sheet(iFile, 1)
	// 创建新的 CSV 文件
	csvFile, _ := os.Create(output)
	defer func() { _ = csvFile.Close() }()
	// 初始化 CSV writer
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()
	// 将 Excel 行数据写入 CSV 文件
	for _, iRow := range iRows {
		_ = csvWriter.Write(iRow)
	}
}

// SaveAsCSVDeduplication 保存为 CSV 文件
func (e *XLS) SaveAsCSVDeduplication(input string, output string) {
	iFile := e.Open(input)
	defer func() { _ = iFile.Close() }()
	_, iRows := e.Sheet(iFile, 1)
	// 创建新的 CSV 文件
	csvFile, _ := os.Create(output)
	defer func() { _ = csvFile.Close() }()
	// 初始化 CSV writer
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()
	// 将 Excel 行数据写入 CSV 文件
	uniqueKeys := make(map[string]bool)
	// 将 Excel 行数据写入 CSV 文件
	for _, iRow := range iRows {
		if len(iRow) == 0 {
			continue
		}

		firstColValue := iRow[0]

		if _, exists := uniqueKeys[firstColValue]; !exists {
			uniqueKeys[firstColValue] = true
			_ = csvWriter.Write(iRow)
		}
	}
}

// Sheets 获取所有 sheet
func (e *XLS) Sheets(f *excelize.File) []string {
	return f.GetSheetList()
}

// Sheet 获取第几个 sheet
func (e *XLS) Sheet(f *excelize.File, sheetNum int) (sheet string, rows [][]string) {
	sheet = e.Sheets(f)[sheetNum-1]
	rows, _ = f.GetRows(sheet)
	return
}

// IntCol 获取第一列
func (e *XLS) IntCol(f *excelize.File, sheetNum, skipLine int) (col []string) {
	_, rows := e.Sheet(f, 1)
	for i, row := range rows {
		if i < skipLine {
			continue
		}
		col = append(col, row[sheetNum-1])
	}
	return
}

// InCol 判断字符串是否在指定列中
func (e *XLS) InCol(value string, col []string) bool {
	for _, v := range col {
		if v == value {
			return true
		}
	}
	return false
}

// IsDuplicate 判断是否重复, 返回是否重复和重复的索引
func (e *XLS) IsDuplicate(rows [][]string, row []string) (bool, int) {
	// targetHash := hashRow(row)
	for i, r := range rows {
		// if hashRow(r) == targetHash {
		// 	return true, i
		// }
		if reflect.DeepEqual(r, row) {
			return true, i
		}
	}
	return false, -1
}

// ColIndex Excel 列名转换为列索引
func (e *XLS) ColIndex(col string) int {
	return colIndex(col)
}

// colIndex Excel 列名转换为列索引
func colIndex(column string) int {
	result := 0
	length := len(column)

	for i := 0; i < length; i++ {
		result *= 26
		result += int(column[i]-'A') + 1
	}

	return result - 1
}

func hashRow(row []string) string {
	hasher := sha256.New()
	for _, cell := range row {
		hasher.Write([]byte(cell))
	}
	hashBytes := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashBytes)
}

// NewWaitGroup 新建 WaitGroup
func NewWaitGroup() *sync.WaitGroup {
	return &sync.WaitGroup{}
}
