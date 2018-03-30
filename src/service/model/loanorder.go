package model

// 发布车源
type LoanOrder struct {
	SRC_IP                   string //源IPv4地址
	DST_IP                   string //目的IPv4地址
	SRC_PORT                 string //源IPv4端口
	DST_PORT                 string //目的IPv4端口
	MAC                      string //客户端MAC地址
	CAPTURE_TIME             string //数据产生时间
	IMSI                     string //国际移动用户标识号
	EQUIPMENT_ID             string //移动设备特征码
	HARDWARE_SIGNATURE       string //硬件特征码
	LONGITUDE                string //数据产生时的经度
	LATITUDE                 string //数据产生时的纬度
	TERMINAL_TYPE            string //终端类型
	TERMINAL_MODEL           string //终端型号
	TERMINAL_OS_TYPE         string //终端操作系统类型
	SOFTWARE_NAME            string //软件名称
	DATA_LAND                string //数据产生时的行政区划
	APPLICATION_TYPE         string //网络应用代码
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
