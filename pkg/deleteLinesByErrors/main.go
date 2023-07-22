package main

import (
	"strconv"

	"log"

	"kit/logger"
	"kit/pkg/xls"
	"kit/tools/timex"
)

const (
	ErrorsTitleRow = 1 // 错误文件标题行
)

func init() {
	logger.InitZap()
}

func main() {

	defer timex.Cost()()
	var (
		e     = xls.New()          // Excel 实例
		iRows [][]string           // 输入文件行
		col   []string             // 错误文件列
		wg    = xls.NewWaitGroup() // 等待组
	)
	wg.Add(3)
	go func() {
		defer wg.Done()
		// 读取输入文件
		// TODO 文件太大，读取时间太长，需要优化
		iFile := e.OpenInput()
		log.Printf("【读取输入文件】打开完成，正在读取...请稍后 \n")
		defer func() { _ = iFile.Close() }()
		_, iRows = e.Sheet(iFile, 1)
		log.Printf("【读取输入文件】读取完成，共 %d 行 \n", len(iRows))
	}()
	go func() {
		defer wg.Done()
		// 初始化输出目录
		e.InitOutputDir()
		// 拷贝输入文件
		e.CopyTemplateToOutput()
	}()
	go func() {
		defer wg.Done()
		// 读取errors文件
		eFile := e.OpenErrors()
		defer func() { _ = eFile.Close() }()
		col = e.IntCol(eFile, 1, ErrorsTitleRow)
		log.Printf("【读取错误文件】读取完成，共 %d 行 \n", len(col))
	}()
	wg.Wait()

	// 读取输出文件
	oFile := e.OpenOutput()
	defer func() { _ = oFile.Close() }()
	oSheet, _ := e.Sheet(oFile, 1)

	var line = 0
	for i, iRow := range iRows {
		if e.InCol(strconv.Itoa(i+1), col) {
			log.Printf("【删除错误行】删除第 %d 行 \n", i+1)
			continue
		}
		line++
		e.WriteToRow(oFile, oSheet, e.NewRow(iRow...), line)
	}
	e.Save(oFile)
}
