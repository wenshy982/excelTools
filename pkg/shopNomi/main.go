package main

import (
	"log"

	"kit/logger"
	"kit/pkg/xls"
	"kit/tools/timex"
	"kit/tools/waitgroup"
)

func init() {
	logger.InitZap()
}

const (
	inputTitleRow = 3 // 输入文件标题行数
	allSites      = "所有站点都报名"
)

func main() {
	defer timex.Cost()()

	var (
		e     = xls.New()       // Excel 实例
		iRows [][]string        // 输入文件行
		wg    = waitgroup.New() // 等待组
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		// 读取输入文件
		iFile := e.OpenInput()
		log.Println("【读取输入文件】打开完成，正在读取...请稍后")
		defer func() { _ = iFile.Close() }()
		_, iRows = e.Sheet(iFile, 1)
		log.Println("【读取输入文件】读取完成，共", len(iRows), "行")
	}()
	go func() {
		defer wg.Done()
		// 初始化输出目录
		e.InitOutputDir()
		// 复制模板
		e.CopyTemplateToOutput()
	}()
	wg.Wait()

	// 读取输出文件
	oFile := e.OpenOutput()
	defer func() { _ = oFile.Close() }()
	oSheet, oRows := e.Sheet(oFile, 1)
	oRowsLen := len(oRows)

	// 存储所有站点都报名的行
	var allSitesShops [][]string
	// 写入单站点报名的行
	for i, row := range iRows {
		// 跳过标题行
		if i < inputTitleRow {
			continue
		}
		// 跳过所有站点都报名的行
		if row[4] == allSites {
			allSitesShops = append(allSitesShops, row)
			continue
		}
		// 写入单站点报名的行
		r := e.NewRow(row[0], row[1], row[2], row[3], row[4], row[5])
		oRowsLen++
		e.WriteToRow(oFile, oSheet, r, oRowsLen)
		log.Println("【单站点报名写入】", r)
	}

	// 遍历所有站点
	for _, site := range xls.SIPSites {
		// 写入所有站点都报名的店铺
		for _, shop := range allSitesShops {
			r := e.NewRow(shop[0], shop[1], shop[2], shop[3], site, shop[5], "", "", "", "", "所有店铺")
			oRowsLen++
			e.WriteToRow(oFile, oSheet, r, oRowsLen)
			log.Println("【所有站点都报名写入】", r)
		}
	}

	e.Save(oFile)
}
