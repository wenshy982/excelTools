package main

type EnumFinalListType int

const (
	EnumFinalListTypeUnknown  EnumFinalListType = iota
	EnumFinalListTypeCFS                        // CFS
	EnumFinalListTypeCampaign                   // Campaign
	EnumFinalListTypeEnd                        // 结束标记
)

// IsValid 判断枚举值是否有效
func (e EnumFinalListType) IsValid() bool {
	return e > EnumFinalListTypeUnknown && e < EnumFinalListTypeEnd
}

// String 枚举值的字符串表示
func (e EnumFinalListType) String() string {
	switch e {
	case EnumFinalListTypeCFS:
		return "CFS"
	case EnumFinalListTypeCampaign:
		return "Campaign"
	default:
		return ""
	}
}
