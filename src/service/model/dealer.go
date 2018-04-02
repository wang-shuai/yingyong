package model

// 商户
type Dealer struct {
	Base

	ACCOUNT_ID string //用户ID
	ACCOUNT string //用户账号
	PRINCIPAL_NAME string //负责人姓名
	BINDING_PHONE string //绑定手机号码
	COMPANY_NAME string //公司名称
	DEALER_TYPE string //经销商类型   1-经纪人	2-4S店  	3-专业公司  	4-厂商  	5-其他 	6-集团
	DEALER_ADDRESS string //经销商地址
}
