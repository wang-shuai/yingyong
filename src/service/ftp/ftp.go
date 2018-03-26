package ftp

import(
	"github.com/jlaffaye/ftp"
	"fmt"
	"os"
)
type FtpOperation struct {
}

func(this * FtpOperation) FtpUploadFile(ftpserver, ftpuser, pw, localFile, remoteSavePath, saveName string) {
	ftpSvrConn, err := ftp.Connect(ftpserver)
	if err != nil {
		fmt.Println("链接远程ftp失败：",err)
	}
	err = ftpSvrConn.Login(ftpuser, pw)
	if err != nil {
		fmt.Println("登陆远程ftp错误：",err)
	}
	////注意是 pub/log，不能带“/”开头
	//ftpSvrConn.ChangeDir(`C:\ftp\xinche`)
	//dir, err := ftpSvrConn.CurrentDir()
	//fmt.Println(dir)

	err = ftpSvrConn.MakeDir(remoteSavePath)
	if err!=nil{
		fmt.Println("创建远程ftp目录失败：",err)
	}
	ftpSvrConn.ChangeDir(remoteSavePath)
	dir, _ := ftpSvrConn.CurrentDir()
	fmt.Println(dir)
	file, err := os.Open(localFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	err = ftpSvrConn.Stor(saveName, file)
	if err != nil {
		fmt.Println(err)
	}
	ftpSvrConn.Logout()
	ftpSvrConn.Quit()
	fmt.Println("success upload file:", localFile)
}