package model

import "os"

const (
	AppType             = 1430015
	XmlIndexName        = "GAB_ZIP_INDEX.xml"
	tempdir      string = "tempdata"
)

const (
	UserCode       = "WA_BASIC_0009"    // 注册用户
	DealerUserCode = "WA_BASIC_0009_04" //店铺注册用户信息
	LoginCode      = "WA_SOURCE_0029"   //登陆日志
	SubscribeCode  = "WA_SOURCE_0078"   //订阅
	CollectionCode = "WA_SOURCE_0065"   // 收藏
	EvaluateCode   = "WA_SOURCE_0115"   // 车辆评估
	UcarCode       = "WA_SOURCE_0114"   // 车源发布
	LoanOrderCode  = "WA_SOURCE_0116"   //贷款购车
	BookCode       = "WA_SOURCE_0117"   //贷款购车
	DealerCode     = "WA_SOURCE_0121"   //商家申请
)

const (
	UserDir       = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "user" + string(os.PathSeparator)       // 注册用户
	DealerUserDir = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "dealeruser" + string(os.PathSeparator) // 注册用户
	LoginDir      = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "login" + string(os.PathSeparator)      //登陆日志
	SubscribeDir  = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "subscribe" + string(os.PathSeparator)  //关注订阅
	CollectionDir = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "collect" + string(os.PathSeparator)    // 浏览
	EvaluateDir   = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "evaluate" + string(os.PathSeparator)   // 车辆评估
	UcarDir       = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "ucar" + string(os.PathSeparator)       // 车源发布
	LoanOrderDir  = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "loanorder" + string(os.PathSeparator)  // 贷款购车
	BookDir       = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "book" + string(os.PathSeparator)       // 预约记录
	DealerDir     = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "dealer" + string(os.PathSeparator)     // 商家申请
	OutputDir     = string(os.PathSeparator) + "output" + string(os.PathSeparator)                                          // 输出文件夹
)
