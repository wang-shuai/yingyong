package model

// 商家
type DealerAccount struct {
	NAME                string //用户姓名
	SEXCODE             string //性别
	CERTIFICATE_TYPE    string //证件类型
	CERTIFICATE_CODE    string //证件号码
	MOBILE              string //手机号码
	REG_ACCOUNT_TYPE    string //网络应用代码
	ACCOUNT_ID          string //用户ID
	REG_ACCOUNT         string //注册帐号
	REGIS_NICKNAME      string //注册昵称
	REGIS_TIME          string //注册时间
	IP_ADDRESS          string //注册IP地址
	PORT                string //注册端口
	MAC_ADDRESS         string //注册MAC地址
	POSTAL_ADDRESS      string //联系地址
	CONTACTOR_TEL       string //联系电话
	BIRTHDAY            string //出生日期
	COMPANY             string //公司名
	SAFE_QUESTION       string //安全问题
	SAFE_ANSWER         string //安全问题答案
	ACTIVITE_TYPE       string //激活类型
	ACTIVITE_ACCOUNT    string //激活用帐号
	PASSWORD            string //密码
	IMEI                string //注册IMEI
	IMSI                string //注册IMSI
	BAIDUMAP            string //百度 经纬度 逗号分割
	LONGITUDE           string //注册经度
	LATITUDE            string //注册纬度
	SITE_ADDRESS        string //注册地址
	ORIGIN_PLACE        string //籍贯
	OFTEN_ADDRESS       string //居住地
	SHOP_ID             string //店铺ID
	SHOP_NAME           string //店铺名称
	DATA_LAND           string //数据产生时的行政区划
	ACCOUNT_ACTION_TYPE string //动作类型
	DVR_CITY            string //收件人所在城市（中文描述）
	DVR_COUNTRY         string //收件人所在国家（中文描述）
	DVR_DEVISIONCODE    string //收件人行政编码（6位）
	DVR_ADRESSDETAIL    string //收件人详细地址
	DVR_POSTCODE        string //收件人邮政编码
	DVR_NAME            string //收件人姓名全称
	DVR_ID              string //收件人地址编号
	DVR_MOBILE          string //收件人手机号
	DVR_TELEPHONE       string //收件人电话号码（归一化）
	DVR_PROVINCE        string //收件人所在省（中文描述）
	DVR_AREA            string //收件人所在区（中文描述）
	DVR_TOWN            string //收件人所在街道（中文描述）
	DVR_STATUS          string //收件人状态
	DVR_TOWN_NAME       string //收件人乡镇信息
	DVR_TOWN_CODE       string //收件人乡镇编码
	DEFAULTADDRESS      string //默认地址
	DVR_CAPTURETIME     string //创建时间（绝对秒数）
	DVR_UPDATETIME      string //修改时间（绝对秒数）
}
