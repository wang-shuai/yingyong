package model

import "os"

const (
	AppType = 1430015
	XmlIndexName = "GAB_ZIP_INDEX.xml"
	tempdir  string = "tempdata"
)

const (
	UserCode   = "WA_BASIC_0009"  // 注册用户
	LoginCode  = "WA_SOURCE_0029" //登陆日志
	FocusCode  = "WA_SOURCE_0078" //关注订阅
	BrowseCode = "WA_SOURCE_0065" // 浏览
)

const (
	UserDir   = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "user" + string(os.PathSeparator)   // 注册用户
	LoginDir  = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "login" + string(os.PathSeparator)  //登陆日志
	FocusDir  = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "focus" + string(os.PathSeparator)  //关注订阅
	BrowseDir = string(os.PathSeparator) + tempdir + string(os.PathSeparator) + "browse" + string(os.PathSeparator) // 浏览
	OutputDir = string(os.PathSeparator) + "output" + string(os.PathSeparator)   // 输出文件夹
)
