package main

import (
	"kit/logger"
	"kit/pkg/xls"
	"kit/tools/timex"
)

func init() {
	logger.InitZap()
}

func main() {
	defer timex.Cost()()
	e := xls.New(xls.WithOutputCSV("." + xls.Sep + "output" + xls.Sep + "A-itemId跑数用.csv"))
	e.InitOutputDir()
	e.SaveAsCSVDeduplication(e.Input, e.OutPutCSV)
}
