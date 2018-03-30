package model

type EvaluateUcar struct {
	SRC_IP                string //源IPv4地址
	DST_IP                string //目的IPv4地址
	SRC_PORT              string //源IPv4端口
	DST_PORT              string //目的IPv4端口
	MAC                   string //客户端MAC地址
	CAPTURE_TIME          string //数据产生时间
	IMSI                  string //国际移动用户标识号
	EQUIPMENT_ID          string //移动设备特征码
	HARDWARE_SIGNATURE    string //硬件特征码
	LONGITUDE             string //数据产生时的经度
	LATITUDE              string //数据产生时的纬度
	TERMINAL_TYPE         string //终端类型
	TERMINAL_MODEL        string //终端型号
	TERMINAL_OS_TYPE      string //终端操作系统类型
	SOFTWARE_NAME         string //软件名称
	DATA_LAND             string //数据产生时的行政区划
	APPLICATION_TYPE      string //网络应用代码
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
