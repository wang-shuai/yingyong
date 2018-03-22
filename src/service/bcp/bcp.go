package bcp

import (
	"os"
	"../index"
	"fmt"
	"../model"
	"io/ioutil"
	"github.com/alexmullins/zip"
	"time"
	"strconv"
	"io"
	"bytes"
)

type BcpOperation struct {
}

const (
	pagesize int64 = 5000
)

func (bo *BcpOperation) ZipUserInfo() {
	user := new(UserBcp)
	filelist, err := user.WriteUserBcp()//写bcp文件
	if err != nil {
		fmt.Println("写入注册用户bcp文件失败：", err)
		return
	}
	idx := new(index.Index)
	idx.BuildUserIdx(filelist)//写索引文件

	filedir := model.Basepath + model.UserDir

	bcpzip(filedir,model.UserCode) //加密打包zip

}

func bcpzip(filedir,code string){
	now := time.Now()
	files, err := ioutil.ReadDir(filedir)
	if err != nil {
		fmt.Println("打开用户bcp临时文件失败：", err)
		return
	}

	//zip path
	var path string
	datepath := fmt.Sprintf("%4d%02d%02d", now.Year(), now.Month(), now.Day())
	timepath := fmt.Sprintf("%02d%02d", now.Hour(), now.Minute())
	zipname := strconv.Itoa(model.AppType) + "-" + strconv.FormatInt(now.Unix(), 10) + "-11-1-00001.zip"

	path = model.Basepath + model.OutputDir + datepath + string(os.PathSeparator) + code + string(os.PathSeparator) + timepath + string(os.PathSeparator)
	if _, err := os.Open(path); err != nil {
		os.MkdirAll(path, os.ModePerm)
	}

	fzip, _ := os.Create(path + zipname)
	zipw  := zip.NewWriter(fzip)
	defer zipw.Close()
	pwd,_:=model.Cfg.String("zippwd")
	for _, file := range files {
		w, err := zipw.Encrypt(file.Name(), pwd)
		if err != nil {
			fmt.Println(err)
			continue
		}
		contents, err := ioutil.ReadFile(filedir + file.Name())
		if err != nil {
			fmt.Println("写入zip包时读取文件失败：", filedir+file.Name(), err)
			continue
		}
		_, err = io.Copy(w, bytes.NewReader(contents))
		if err != nil {
			fmt.Println("写入zip包时copy文件数据流失败：", filedir+file.Name(), err)
			continue
		}
	}
	zipw.Flush()
}
