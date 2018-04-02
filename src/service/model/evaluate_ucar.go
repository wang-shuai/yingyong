package model

// 车辆评估
type EvaluateUcar struct {
	Base

	ACCOUNT_ID            string //用户ID
	ACCOUNT               string //用户账号
	CAR_BRAND             string //品牌
	CAR_TYPE              string //车型
	CARD_TIME             string //上牌时间
	SALE_CITY             string //出售城市
	MILEAGE               string //行驶里程
	VEHICLE_CONDITION     string //车况
	LICENSE_PLATE_SITE    string //牌照地
	USED_CAR_PRICE        string //二手车标注价格
	TRANSFER_NUMBER       string //过户次数
	EXPECTED_SELLING_TIME string //预期售出时间
}
