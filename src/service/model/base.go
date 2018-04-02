package model

type Base struct {
	SRC_IP string //源IPv4地址
	DST_IP string //目的IPv4地址
	SRC_PORT string //源IPv4端口
	DST_PORT string //目的IPv4端口
	MAC string //客户端MAC地址
	CAPTURE_TIME string //数据产生时间
	IMSI string //国际移动用户标识号
	EQUIPMENT_ID string //移动设备特征码
	HARDWARE_SIGNATURE string //硬件特征码
	LONGITUDE string //数据产生时的经度
	LATITUDE string //数据产生时的纬度
	TERMINAL_TYPE string //终端类型
	TERMINAL_MODEL string //终端型号
	TERMINAL_OS_TYPE string //终端操作系统类型
	SOFTWARE_NAME string //软件名称
	DATA_LAND string //数据产生时的行政区划
	APPLICATION_TYPE string //网络应用代码
}