package main

import (
	"golang-services/jingyong/bcp"
	"sync"
	"golang-services/jingyong/ftp"
	"golang-services/jingyong/model"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		bcp.ZipUserInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bcp.ZipDealerUserInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bcp.ZipCollectionInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bcp.ZipSubscribeInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bcp.ZipUcar()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bcp.ZipEvaluate()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bcp.ZipLoanOrder()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bcp.ZipBooks()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bcp.ZipDealers()
	}()

	wg.Wait()

	//上传ftp
	ftpOp := new(ftp.FtpOperation)
	server, _ := model.Cfg.String("ftp.server")
	port, _ := model.Cfg.String("ftp.port")
	username, _ := model.Cfg.String("ftp.username")
	password, _ := model.Cfg.String("ftp.password")

	localpath := model.Basepath + model.OutputDir
	ftpOp.UploadDir(server, port, username, password, localpath, "/")
}
