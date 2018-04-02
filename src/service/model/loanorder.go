package model

// 发布车源
type LoanOrder struct {
	Base

	ACCOUNT_ID               string //用户ID
	ACCOUNT                  string //用户账号
	BUY_CITY                 string //购车城市
	Name                     string //中文姓名
	SFZH                     string //身份证号码
	BANK_ACCOUNT_NUM         string //银行卡号
	RELATIONSHIP_MOBILEPHONE string //关系人手机号码
	CAREER_STYLE             string //职业类别
	DRIVING_LICENSE          string //驾照号
	LOAN                     string //贷款金额
	CAR_BRAND                string //品牌
	CAR_TYPE                 string //车型
	CARD_TIME                string //上牌时间
	MILEAGE                  string //行驶里程
	VEHICLE_CONDITION        string //车况
	LICENSE_PLATE_SITE       string //牌照地
	USED_CAR_PRICE           string //二手车价格
}
