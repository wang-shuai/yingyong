package model

// 发布车源
type UcarBaseInfo struct {
	Base

	ACCOUNT_ID         string //用户ID
	ACCOUNT            string //用户账号
	BINDING_PHONE      string //绑定手机号码
	USED_CAR_NAME      string //二手车名称
	USED_CAR_PRICE     string //二手车价格
	USED_CAR_URL       string //二手车链接
	USED_CAR_ID        string //车辆ID
	CAR_BRAND          string //品牌
	CAR_TYPE           string //车型
	CARD_TIME          string //上牌时间
	SALE_CITY          string //出售城市
	MILEAGE            string //行驶里程
	VEHICLE_CONDITION  string //车辆照片实体文件列表
	LICENSE_PLATE_SITE string //牌照地
}
