package main

import (
	"fmt"
	"log"
	"sync"

	"kit/logger"
	"kit/pkg/xls"
	"kit/tools/timex"
)

func init() {
	logger.InitZap()
}

const (
	templateRowLen int = 3 // 模板文件行数
	column         int = 6 // 根据第几列分表
)

func main() {
	defer timex.Cost()()
	var e = xls.New()
	// 初始化输出目录
	e.InitOutputDir()
	// 读取输入文件
	iFile := e.OpenInput()
	defer func() { _ = iFile.Close() }()
	_, iRows := e.Sheet(iFile, 1)
	fmt.Printf("【读取输入文件】读取完毕，共 %d 行数据 \n", len(iRows))

	var mapFile = make(map[string][][]string)
	// 创建文件
	var wg sync.WaitGroup
	var noUseLine, createTimes = 0, 0
	for i, row := range iRows {
		if row == nil {
			continue
		}
		c := row[column]
		if c == "" || i < templateRowLen {
			noUseLine++
			continue
		}
		// 判断是否已经创建
		_, ok := mapFile[c]
		if ok {
			mapFile[c] = append(mapFile[c], row)
			continue
		}
		mapFile[c] = append(mapFile[c], row)
		wg.Add(1)
		// 并发创建
		go func(i int, c string) {
			defer wg.Done()
			createTimes++
			e.CopyTemplateTo(e.OutputDir + xls.Sep + c + ".xlsx")
		}(i, c)
	}
	wg.Wait()
	log.Printf("【创建文件】共 %d 行数据，其中 %d 行为空行或标题行 \n", len(iRows), noUseLine)
	log.Printf("【创建文件】共创建%d次 \n", createTimes)
	log.Printf("【创建文件】创建完毕，共 %d 个文件 \n", len(mapFile))

	// 写入数据
	var job = make(chan int, 20)
	for key, value := range mapFile {
		dst := e.OutputDir + xls.Sep + key + ".xlsx"
		wg.Add(1)
		job <- 1
		go func(key, dst string, rows [][]string) {
			defer wg.Done()
			defer func() { <-job }()
			f := e.Open(dst)
			defer func() { _ = f.Close() }()
			fSheet, fRows := e.Sheet(f, 1)
			fRowsLen := len(fRows)
			for _, row := range rows {
				fRowsLen++
				e.WriteToRow(f, fSheet, e.NewRow(row...), fRowsLen)
			}
			e.Save(f)
		}(key, dst, value)
		fmt.Printf("【写入数据】%s 写入完毕\n", key)
	}
	wg.Wait()
}
