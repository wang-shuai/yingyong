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
	Postal_Address   string //	联系地址
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

type DealerAccount struct {
	NAME             string //用户姓名
	SEXCODE          string //性别
	CERTIFICATE_TYPE string //证件类型
	CERTIFICATE_CODE string //证件号码
	MOBILE           string //手机号码
	REG_ACCOUNT_TYPE string //网络应用代码
	ACCOUNT_ID       string //用户ID
	REG_ACCOUNT      string //注册帐号
	REGIS_NICKNAME   string //注册昵称
	REGIS_TIME       string //注册时间
	IP_ADDRESS       string //注册IP地址
	PORT             string //注册端口
	MAC_ADDRESS      string //注册MAC地址
	POSTAL_ADDRESS   string //联系地址
	CONTACTOR_TEL    string //联系电话
	BIRTHDAY         string //出生日期
	COMPANY          string //公司名
	SAFE_QUESTION    string //安全问题
	SAFE_ANSWER      string //安全问题答案
	ACTIVITE_TYPE    string //激活类型
	ACTIVITE_ACCOUNT string //激活用帐号
	PASSWORD         string //密码
	IMEI             string //注册IMEI
	IMSI             string //注册IMSI
	BAIDUMAP         string //百度 经纬度 逗号分割
	LONGITUDE        string //注册经度
	LATITUDE         string //注册纬度
	SITE_ADDRESS     string //注册地址
	ORIGIN_PLACE     string //籍贯
	OFTEN_ADDRESS    string //居住地
	SHOP_ID          string //店铺ID
	SHOP_NAME        string //店铺名称
	DATA_LAND        string //数据产生时的行政区划
}

type Collection struct {
	SRC_IP                  string //源IPv4地址
	DST_IP                  string //目的IPv4地址
	SRC_PORT                string //源IPv4端口
	DST_PORT                string //目的IPv4端口
	MAC                     string //客户端MAC地址
	CAPTURE_TIME            string //数据产生时间
	IMSI                    string //国际移动用户标识号
	EQUIPMENT_ID            string //移动设备特征码
	HARDWARE_SIGNATURE      string //硬件特征码
	LONGITUDE               string //数据产生时的经度
	LATITUDE                string //数据产生时的纬度
	TERMINAL_TYPE           string //终端类型
	TERMINAL_MODEL          string //终端型号
	TERMINAL_OS_TYPE        string //终端操作系统类型
	SOFTWARE_NAME           string //软件名称
	DATA_LAND               string //数据产生时的行政区划
	APPLICATION_TYPE        string //网络应用代码
	USER_INTENRALID         string //用户ID
	USER_ACCOUNT            string //用户账号
	NEWS_ID                 string //新闻/信息/消息/ID
	LABEL                   string //收藏标签（分类）
	LIKE_TYPE               string //点赞类型（支持/反对）
	ACTION_TIME             string //时间
	FILE_MD5                string //文件特征值MD5
	FILE_ID                 string //文件ID
	FILE_URL                string //文件路径/URL
	ACTION_TYPE             string //操作类型（收藏/点赞）
	GOODS_ID                string //浏览商品ID
	GOODS_NAME              string //新闻信息消息标题/商品名称
	GOODS_COMMENT           string //商品评论
	GOODS_PRICE             string //商品价格
	SHOP_ID                 string //商家ID
	SHOP_NAME               string //商家名称
	COLLECTPOSITION_NAME    string //收藏位置名称（家、公司等）
	COLLECTPOSITION_ADDRESS string //收藏位置详细地址
}

type Subscribe struct {
	SRC_IP              string //源IPv4地址
	DST_IP              string //目的IPv4地址
	SRC_PORT            string //源IPv4端口
	DST_PORT            string //目的IPv4端口
	MAC                 string //客户端MAC地址
	CAPTURE_TIME        string //数据产生时间
	IMSI                string //国际移动用户标识号
	EQUIPMENT_ID        string //移动设备特征码
	HARDWARE_SIGNATURE  string //硬件特征码
	LONGITUDE           string //数据产生时的经度
	LATITUDE            string //数据产生时的纬度
	TERMINAL_TYPE       string //终端类型
	TERMINAL_MODEL      string //终端型号
	TERMINAL_OS_TYPE    string //终端操作系统类型
	SOFTWARE_NAME       string //软件名称
	DATA_LAND           string //数据产生时的行政区划
	APPLICATION_TYPE         string // string //网络应用代码
	IDENTIFICATION_TYPE string //合作账号类型
	A_IDEN_ID           string //用户ID
	A_IDEN_STRING       string //用户账号
	A_PHONE             string //手机号
	ACTION_TYPE         string //操作类型（关注/订阅/取消关注/取消订阅）
	UPDATE_time         string //操作时间
	B_IDEN_ID           string //被关注/订阅用户ID
	SUB_NAME            string //关注/订阅名称（文件/频道）
	SUB_ID              string //关注/订阅（文件/频道）ID
	SUB_NUM             string //关注/订阅数量
	FILE_SIZE           string //关注/订阅文件大小
	MAINFILE            string //关注/订阅文件实体
}
