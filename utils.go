package adfunds


//结算中心结算数据
type SettleMsg struct {
	//操作类型 奖励reward，联盟提成commission，流量主地址异常回收 abnormal,转账transfer
	ActionType string `json:"actionType"`
	//操作地址
	Address string `json:"address"`
	//操作金额
	Amount string `json:"amount"`
	//余额
	Balance string `json:"balance"`
	//广告hash 只在给流量主结算adcoin 即ActionType为reward时不为空其他情况为空
	AdHash string `json:"adHash"`
}

//账户转账数据
type TransferMsg struct {
	//操z`作类型 transfer 转账
	ActionType string `json:"actionType"`
	//转出账户
	AddressOut string `json:"addressOut"`
	//转入账户
	AddressIn string `json:"addressIn"`
	//转账金额金额
	Amount string `json:"amount"`
	//转出账户余额
	AddressOutBalance string `json:"addressOutBalance"`
	//转入账户余额
	AddressInBalance string `json:"addressInBalance"`

}

type Address struct {
	Address string `json:"address"`
	Exists bool `json:"exists"`
	Balance string `json:"balance"`
}


//浮点数计算方法
func StrToFloat64Round(str string, prec int, round bool) (float64,error) {
	f,err := strconv.ParseFloat(str,64)
	return Precision(f,prec,round),err
}
func Precision(f float64, prec int, round bool) float64 {
	pow10_n := math.Pow10(prec)
	if round {
		return math.Trunc((f + 0.5/pow10_n)*pow10_n) / pow10_n
	}
	return math.Trunc((f)*pow10_n) / pow10_n
}

func FloatToStrWithPrec(fv float64, prec int) string {
	return strconv.FormatFloat(fv, 'f', prec, 64)
}

func FloatToStr(fv float64) string {
	return strconv.FormatFloat(fv, 'f', 2, 64)
}
