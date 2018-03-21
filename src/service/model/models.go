package model

type User struct {
	Name             string //	用户姓名
	SexCode          string //	性别
	Certificate_Type string //	证件类型
	Certificate_Code string //	证件号码
	Mobile           string //	手机号码
	Reg_Account_Type string //	网络应用代码
	Account_Id       string //	用户ID
	Reg_Account      string //	注册帐号
	Regis_NickName   string //	注册昵称
	Regis_Time       string //	注册时间
	Ip_Address       string //	注册IP地址
	Port             string //	注册端口
	Mac_Address      string //	注册MAC地址
	Postal_ddress    string //	联系地址
	Contactor_Tel    string //	联系电话
	Birthday         string //	出生日期
	Company          string //	公司名
	Safe_Question    string //	安全问题
	Safe_Answer      string //	安全问题答案
	Activite_Type    string //	激活类型
	Activite_Account string //	激活用帐号
	Password         string //	密码
	IMEI             string //	注册IMEI
	IMSI             string //	注册IMSI
	Longitude        string //	注册经度
	Latitude         string //	注册纬度
	Site_Address     string //	注册地址
	Origin_Place     string //	籍贯
	Often_Address    string //	居住地
	Data_Land        string //	数据产生时的行政区划

}
