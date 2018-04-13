package ftp

import (
	"github.com/jlaffaye/ftp"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
	"os"
	"fmt"
	"strings"
	"path"

	"log"
	"io/ioutil"
)

type FtpOperation struct {
}

func (this *FtpOperation) UploadFile(ftpserver, ftpuser, pw, localFile, remoteSavePath, saveName string) {
	ftpSvrConn, err := ftp.Connect(ftpserver)
	if err != nil {
		flog.Errorf("链接远程ftp失败：%v \n", err)
	}
	err = ftpSvrConn.Login(ftpuser, pw)
	if err != nil {
		flog.Errorf("登陆远程ftp错误：%v \n", err)
	}
	////注意是 pub/log，不能带“/”开头
	//ftpSvrConn.ChangeDir(`C:\ftp\xinche`)
	//dir, err := ftpSvrConn.CurrentDir()
	//flog.Errorf(dir)

	// 如果有 则直接转过去 否则就是没有目录，则创建
	if err = ftpSvrConn.ChangeDir(remoteSavePath); err != nil {

		err = ftpSvrConn.MakeDir(remoteSavePath)
		if err != nil {
			flog.Errorf("创建远程ftp目录失败：%v \n", err)
		}
	}

	ftpSvrConn.ChangeDir(remoteSavePath)
	dir, _ := ftpSvrConn.CurrentDir()
	flog.Errorf("当前ftp服务器目录：%s", dir)
	file, err := os.Open(localFile)
	if err != nil {
		flog.Errorf("打开本地文件%s失败：%v \n", localFile, err)
		fmt.Println("打开本地文件失败：", localFile, err)
	}
	defer file.Close()
	err = ftpSvrConn.Stor(saveName, file)
	if err != nil {
		flog.Errorf("存储ftp文件%s失败：%v \n", saveName, err)
	}
	ftpSvrConn.Logout()
	ftpSvrConn.Quit()
	flog.Infof("success upload file: %s \n", localFile)
}

func (this *FtpOperation) UploadFile_1(ftpserver, ftpuser, pw, localFile, remoteSavePath, saveName string) {
	ftpSvrConn, err := ftp.Connect(ftpserver)
	if err != nil {
		flog.Errorf("链接远程ftp失败：%v \n", err)
	}
	err = ftpSvrConn.Login(ftpuser, pw)
	if err != nil {
		flog.Errorf("登陆远程ftp错误：%v \n", err)
	}

	makealldir(ftpSvrConn, remoteSavePath)

	fmt.Println("请求ftp路径为", remoteSavePath)
	err = ftpSvrConn.ChangeDir(remoteSavePath)
	if err != nil {
		flog.Errorf("创建ftp路径%s失败 \n", remoteSavePath)
		fmt.Printf("创建ftp路径%s失败\n", remoteSavePath)
	}
	dir, _ := ftpSvrConn.CurrentDir()
	flog.Infof("当前ftp服务器目录：%s\n", dir)
	fmt.Printf("当前ftp服务器目录：%s \n", dir)
	file, err := os.Open(localFile)
	if err != nil {
		flog.Errorf("打开本地文件%s失败：%v \n", localFile, err)
		fmt.Println("打开本地文件失败：", localFile, err)
	}
	defer file.Close()
	err = ftpSvrConn.Stor(saveName, file)
	if err != nil {
		flog.Errorf("存储ftp文件%s失败：%v \n", saveName, err)
		fmt.Println("存储ftp文件%s失败：%v \n", saveName, err)
	}
	ftpSvrConn.Logout()
	ftpSvrConn.Quit()
	flog.Infof("success upload file: %s \n", localFile)
}

func makealldir(ftpSvrConn *ftp.ServerConn, remoteSavePath string) {
	// 如果有 则直接转过去 否则就是没有目录，则创建
	if err := ftpSvrConn.ChangeDir(remoteSavePath); err != nil {
		paths := strings.Split(remoteSavePath, string(os.PathSeparator))

		dest := ""
		for _, p := range paths {
			if len(p) > 0 {
				dest = path.Join(dest, p)
				chkormkdir(ftpSvrConn, dest)
			}
		}
	}
}

func chkormkdir(ftpSvrConn *ftp.ServerConn, remoteSavePath string) {
	// 如果有 则直接转过去 否则就是没有目录，则创建
	if err := ftpSvrConn.ChangeDir(remoteSavePath); err != nil {
		err = ftpSvrConn.MakeDir(remoteSavePath)
		if err != nil {
			flog.Errorf("创建远程ftp目录失败：%v %s\n", err, remoteSavePath)
		}
	}
	ftpSvrConn.ChangeDir("/") //回到跟目录
}


func (this *FtpOperation) UploadDir(host, port, user, pw, localPath, remotePath string) {
	sftpClient, _ := connect(user, pw, host, port)
	uploadDirectory(sftpClient, localPath, remotePath)
}

func connect(user, password, host , port string) (*ftp.ServerConn, error) {
	server :=strings.Join([]string{host, port}, ":")
	ftpSvrConn, err := ftp.Connect(server)
	if err != nil {
		flog.Errorf("链接远程ftp失败：%v \n", err)
		fmt.Printf("链接远程ftp失败：%v \n", err)
	}
	err = ftpSvrConn.Login(user, password)
	if err != nil {
		flog.Errorf("登陆远程ftp错误：%v \n", err)
		fmt.Printf("登陆远程ftp错误：%v \n", err)
	}
	return ftpSvrConn,err
}

func uploadFile(ftpSvrConn *ftp.ServerConn, localFilePath string, remoteSavePath string) {
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("os.Open error : ", localFilePath)
		flog.Fatalf("上传ftp打开本地文件异常：%s %v", localFilePath, err)
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)

	fmt.Println("请求ftp路径为", remoteSavePath)
	err = ftpSvrConn.ChangeDir(remoteSavePath)
	if err != nil {
		flog.Errorf("创建ftp路径%s失败 \n", remoteSavePath)
		fmt.Printf("创建ftp路径%s失败\n", remoteSavePath)
	}
	dir, _ := ftpSvrConn.CurrentDir()
	flog.Infof("当前ftp服务器目录：%s\n", dir)
	fmt.Printf("当前ftp服务器目录：%s \n", dir)
	file, err := os.Open(localFilePath)
	if err != nil {
		flog.Errorf("打开本地文件%s失败：%v \n", localFilePath, err)
		fmt.Println("打开本地文件失败：", localFilePath, err)
	}
	defer file.Close()
	err = ftpSvrConn.Stor(path.Join(remoteSavePath, remoteFileName), file)
	if err != nil {
		flog.Errorf("存储ftp文件%s失败：%v \n", remoteFileName, err)
		fmt.Printf("存储ftp文件%s失败：%v \n", remoteFileName, err)
	}
	fmt.Println(localFilePath + "  copy file to remote server finished!")
}

func uploadDirectory(sftpClient *ftp.ServerConn, localPath string, remotePath string) {
	localFiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal("read dir list fail ", err)
	}

	for _, backupDir := range localFiles {
		localFilePath := path.Join(localPath, backupDir.Name())
		remoteFilePath := path.Join(remotePath, backupDir.Name())
		if backupDir.IsDir() {
			sftpClient.MakeDir(remoteFilePath)
			uploadDirectory(sftpClient, localFilePath, remoteFilePath)
		} else {
			uploadFile(sftpClient, path.Join(localPath, backupDir.Name()), remotePath)
		}
	}

	fmt.Println(localPath + "  copy directory to remote server finished!")
}
