package main

// SellerType 输入表 卖家类型/店铺类型
func SellerType(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "店铺类型"
	case EnumFinalListTypeCampaign:
		return "Seller Type"
	default:
		return ""
	}
}

// MstShopID 输入表 MstShopID 列 ==> 模板表 C 列
func MstShopID(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "Mstshopid"
	case EnumFinalListTypeCampaign:
		return "mst_shopid"
	default:
		return ""
	}
}

// MstItemID 输入表 MstItemID 列 ==> 模板表 E 列
func MstItemID(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "Mst itemid"
	case EnumFinalListTypeCampaign:
		return "mst_itemid"
	default:
		return ""
	}
}

// Limit 输入表 Limit 列 ==> 模板表 F 列
func Limit(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "purchase_lmt"
	case EnumFinalListTypeCampaign:
		return "Limit"
	default:
		return ""
	}
}

// Stock 活动库存 输入表 Stock 列 ==> 模板表 G 列
func Stock(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "promo_stock"
	case EnumFinalListTypeCampaign:
		return "Promo stock"
	default:
		return ""
	}
}

// GPAccountName 输入表 GPAccountName 列 ==> 模板表 H 列
func GPAccountName(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "br_gp_account_name"
	case EnumFinalListTypeCampaign:
		return "GP account name"
	default:
		return ""
	}
}

// MstUrl 输入表 MstUrl 列 ==> 模板表 I 列
func MstUrl(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "Mst URL"
	case EnumFinalListTypeCampaign:
		return "mst_URL"
	default:
		return ""
	}
}

// SIPSellerSettlementPrice 输入表 SIPSellerSettlementPrice 列 ==> 模板表 M 列 roundDown 保留两位小数
func SIPSellerSettlementPrice(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "SIP产品\n活动结算价"
	case EnumFinalListTypeCampaign:
		return "SIP卖家结算价"
	default:
		return ""
	}
}

// SIPSellerSettlementCurrency 输入表 SIPSellerSettlementCurrency 列 ==> 模板表 L 列
func SIPSellerSettlementCurrency(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "结算价货币"
	case EnumFinalListTypeCampaign:
		return "SIP卖家结算货币"
	default:
		return ""
	}
}

// Region 输入表 Region 列 ==> 模板表 N 列
func Region(t EnumFinalListType) string {
	switch t {
	case EnumFinalListTypeCFS:
		return "店铺类型"
	case EnumFinalListTypeCampaign:
		return "Region"
	default:
		return ""
	}
}

// InputTitleRow 输入表标题行
func InputTitleRow(t EnumFinalListType) int {
	switch t {
	case EnumFinalListTypeCFS:
		return 2
	case EnumFinalListTypeCampaign:
		return 3
	default:
		return -1
	}
}
