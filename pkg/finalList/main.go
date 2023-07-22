package main

import (
	"fmt"

	"strconv"

	"log"

	"kit/logger"
	"kit/pkg/xls"
	"kit/tools/strx"
	"kit/tools/timex"
)

func init() {
	logger.InitZap()
}

var (
	FinalListType EnumFinalListType // 1: CFS, 2: Campaign
)

const (
	CFSCampaignStartTime = "BR Start Time"     // CFS活动开始时间
	CFSCampaignEndTime   = "BR End Time"       // CFS活动结束时间
	CampaignStartTime    = "4/6/2023 0:00:00"  // Campaign活动开始时间
	CampaignEndTime      = "4/9/2023 23:59:59" // Campaign活动结束时间

	MstUsername = "Mst_username" // Mst Username 列 ==> 模板表 D 列 CFS 才有

	SIP = "SIP" // SIP 列 ==> 模板表 J 列 固定值
	BR  = "BR"  // BR 列 ==> 模板表 K 列 固定值

	// inputTitleRow = 3 // 输入文件标题行数
	oS1TitleRow = 1 // 输出文件Sheet1标题行数
)

func NewTemplateRow(t EnumFinalListType) []string {
	return []string{
		"",                             // A 列 Campaign Start Time
		"",                             // B 列 Campaign End Time
		MstShopID(t),                   // C 列 Mst Shop ID
		"",                             // D 列 mst Shop Username
		MstItemID(t),                   // E 列 Mst Item ID
		Limit(t),                       // F 列 Limit
		"",                             // G 列 Campaign Stock
		GPAccountName(t),               // H 列 GPAccountName
		MstUrl(t),                      // I 列 mst_URL
		"",                             // J 列 SIP 固定值
		"",                             // K 列 BR 固定值
		SIPSellerSettlementPrice(t),    // L 列 SIP卖家结算价 roundDown 保留两位小数
		SIPSellerSettlementCurrency(t), // M 列 SIP卖家结算货币
		Region(t),                      // N 列 Region
	}
}

func main() {
	defer timex.Cost()()

	var (
		e               = xls.New()          // Excel 实例
		iRows           [][]string           // 输入文件行
		wg              = xls.NewWaitGroup() // 等待组
		inputTitleRow   int                  // 输入文件标题行数
		iTitle          []string             // 输入文件标题行
		sellerTypeIndex int                  // Seller Type 列索引
	)

	breakFlag := false
	for {
		// 用户输入 FinalListType
		fmt.Printf("请输入 FinalListType (1: CFS, 2: Campaign): ")
		_, _ = fmt.Scanln(&FinalListType)
		if !FinalListType.IsValid() {
			log.Printf("【输入 FinalListType】失败，FinalListType: %d 无效 1: CFS, 2: Campaign \n", FinalListType)
			continue
		} else {
			log.Printf("【输入 FinalListType】成功, finalList 类型: %s \n", FinalListType.String())
			wg.Add(2)
			go func() {
				defer wg.Done()
				// 读取输入文件
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
			wg.Wait()
			inputTitleRow = InputTitleRow(FinalListType)
			iTitle = iRows[inputTitleRow-1]

			sellerTypeTitle := SellerType(FinalListType)
			sellerTypeIndex = xls.IndexOf(iTitle, sellerTypeTitle, 0)
			if sellerTypeIndex != -1 {
				breakFlag = true
			} else {
				log.Printf("【读取输入文件】读取失败，未找到 %s 列, 请检查输入的 FinalListType（%s）是否正确 \n",
					sellerTypeTitle, FinalListType.String())
			}
		}
		if breakFlag {
			break
		}
	}

	// 读取输出文件
	oFile := e.OpenOutput()
	defer func() { _ = oFile.Close() }()
	oS1, oS1Rows := e.Sheet(oFile, 1)
	oS1RowsLen := len(oS1Rows)

	oRows := make([][]string, 0)
	oRows = append(oRows, oS1Rows...)

	for i, iRow := range iRows {
		// 跳过标题行
		if i < inputTitleRow {
			continue
		}
		// 跳过非SIP站点
		if !xls.InSIP(iRow[sellerTypeIndex]) {
			continue
		}

		var row = NewTemplateRow(FinalListType)
		// 从输入文件读取
		row, err := xls.GetFinalListRowByTitle(iTitle, iRow, row)
		if err != nil {
			log.Fatalf("【读取输入文件】读取失败，错误：%v \n", err)
			return
		}

		// 处理 Campaign Start Time, Campaign End Time
		if FinalListType == EnumFinalListTypeCampaign {
			row[0] = CampaignStartTime
			row[1] = CampaignEndTime
		} else {
			row[0] = iRow[xls.IndexOf(iTitle, CFSCampaignStartTime, 0)]
			row[1] = iRow[xls.IndexOf(iTitle, CFSCampaignEndTime, 0)]
			row[e.ColIndex("D")] = iRow[xls.IndexOf(iTitle, MstUsername, 0)]
		}

		row[e.ColIndex("J")] = SIP
		row[e.ColIndex("K")] = BR
		// L 列 SIP卖家结算价 roundDown 保留两位小数
		floatVal := strx.ToFloat64(row[e.ColIndex("L")])
		row[e.ColIndex("L")] = strconv.FormatFloat(floatVal, 'f', 2, 64)

		// 判断是否重复
		dFlag, _ := e.IsDuplicate(oRows, row)
		if dFlag {
			log.Printf("【写入输出文件】第 %d 行重复，跳过写入 \n", i+1)
			continue
		}
		log.Printf("【写入输出文件】正在写入第 %d 行，内容：%v \n", i+1, row)

		oRows = append(oRows, row)
		oS1RowsLen++
		e.WriteToRow(
			oFile,            // 写入文件
			oS1,              // 写入表
			e.NewRow(row...), // 写入行
			oS1RowsLen,       // 写入行数
		)
	}
	e.Save(oFile)
	log.Printf("【写入输出文件】写入完成，共 %d 行 \n", oS1RowsLen-oS1TitleRow)

	// 读取输出文件
	oFile = e.OpenOutput()
	defer func() { _ = oFile.Close() }()
	oS1, oS1Rows = e.Sheet(oFile, 1)
	oS1RowsLen = len(oS1Rows)

	for oIndex, oRow := range oS1Rows {
		var stock = 0
		itemID := oRow[e.ColIndex("E")]
		for iIndex, iRow := range iRows {
			// 跳过标题行
			if iIndex < inputTitleRow {
				continue
			}
			// 跳过非SIP站点
			if !xls.InSIP(iRow[sellerTypeIndex]) {
				continue
			}
			// 累加ItemID相同的stock
			iItemID := iRow[xls.IndexOf(iTitle, MstItemID(FinalListType), 0)]
			if iItemID != itemID {
				continue
			}

			s := iRow[xls.IndexOf(iTitle, Stock(FinalListType), 0)]
			if FinalListType == EnumFinalListTypeCampaign {
				stock += strx.ToInt(s)
			} else {
				// CFS 还要根据时间段计算stock
				campaignStartTime := iRow[xls.IndexOf(iTitle, CFSCampaignStartTime, 0)]
				oCampaignStartTime := oRow[e.ColIndex("A")]
				if campaignStartTime != oCampaignStartTime {
					continue
				}
				stock += strx.ToInt(s)
			}
		}
		// 跳过标题行
		if oIndex < oS1TitleRow {
			continue
		}
		_ = oFile.SetCellValue(oS1, fmt.Sprintf("G%d", oIndex+1), stock)
		log.Printf("【%s计算stock】正在写入第 G%d 单元格，stock：%v \n", FinalListType.String(), oIndex+1, strconv.Itoa(stock))
	}
	e.Save(oFile)
	log.Printf("【%s计算stock】写入完成，共 %d 行 \n", FinalListType.String(), oS1RowsLen-oS1TitleRow)
}
