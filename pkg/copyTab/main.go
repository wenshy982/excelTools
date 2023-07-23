package main

import (
	"log"

	"kit/logger"
	"kit/pkg/xls"
	"kit/tools/timex"
)

func init() {
	logger.InitZap()
}

const (
	TitleRow = 1 // 表头
)

func main() {
	defer timex.Cost()()
	var (
		e     = xls.New() // Excel 实例
		iRows [][]string  // 输入文件行
	)

	e.InitOutputDir()
	e.CopyInputToOutput()

	// 读取输出文件
	oFile := e.OpenOutput()
	defer func() { _ = oFile.Close() }()
	log.Printf("【读取输出文件】打开完成，正在读取...请稍后 \n")
	sheets := e.Sheets(oFile)
	_, iRows = e.Sheet(oFile, 1)
	log.Printf("【读取输出文件】读取完成，共 %d 个sheet \n", len(sheets))
	log.Printf("【读取输出文件】sheets: %v \n", sheets)
	for i, s := range sheets {
		if i == 0 {
			continue
		}
		e.WriteToRow(oFile, s, e.NewRow(iRows[0]...), 1)
		e.Save(oFile)
	}
	for _, s := range xls.SIPSites {
		log.Printf("【处理输出文件】正在处理 %s 站点 \n", s)
		var line = 1
		for i, iRow := range iRows {
			if i == 0 {
				continue
			}
			col := iRows[i][3]
			if col != s {
				continue
			}
			line++
			e.WriteToRow(oFile, col, e.NewRow(iRow...), line)
			if line%10000 == 0 { // 每 10000 行保存一次
				e.Save(oFile) // 保存, 防止内存溢出
				log.Printf("【处理输出文件】正在处理站点 %s , 已处理 %d 行 \n", s, line)
			}
		}
		e.Save(oFile)
		log.Printf("【处理输出文件】站点 %s 处理完成，共 %d 行 \n", s, line)
	}
}
