package model

// 收藏
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
