package bcp

import (
	"os"
	"golang-services/jingyong/index"
	"fmt"
	"golang-services/jingyong/model"
	"io/ioutil"
	"github.com/alexmullins/zip"
	"time"
	"strconv"
	"io"
	"bytes"
	"golang-services/jingyong/ftp"
	"strings"
	"golang-services/jingyong/tool"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
)

const (
	pagesize int64 = 5000
)

var (
	ftpOp = new(ftp.FtpOperation)
)

func ZipUserInfo() {

	filedir := model.Basepath + model.UserDir
	clean(filedir)

	user := new(UserBcp)
	filelist, err := user.WriteUserBcp() //写bcp文件
	if err != nil {
		flog.Errorf("写入注册用户bcp文件失败：%v \n", err)
		return
	}
	idx := new(index.Index)
	idx.BuildUserIdx(filelist) //写索引文件

	bcpzip(filedir, model.UserCode) //加密打包zip
}

func ZipDealerUserInfo() {
	//model.Basepath + model.OutputDir +

	filedir := model.Basepath + model.DealerUserDir
	clean(filedir)
	DealerUser := new(DealerUserBcp)
	filelist, err := DealerUser.WriteDealerUserBcp() //写bcp文件
	if err != nil {
		flog.Errorf("写入注册商户bcp文件失败：%v \n", err)
		return
	}
	idx := new(index.Index)
	idx.BuildDealerUserIdx(filelist) //写索引文件

	bcpzip(filedir, model.DealerUserCode) //加密打包zip
}

func ZipCollectionInfo() {

	filedir := model.Basepath + model.CollectionDir
	clean(filedir)
	collect := new(CollectionBcp)
	filelist, err := collect.WriteCollectionBcp() //写bcp文件
	if err != nil {
		flog.Errorf("写入收藏bcp文件失败：%v \n", err)
		return
	}
	idx := new(index.Index)
	idx.BuildCollectionIdx(filelist) //写索引文件

	bcpzip(filedir, model.CollectionCode) //加密打包zip
}

func ZipSubscribeInfo() {

	filedir := model.Basepath + model.SubscribeDir
	clean(filedir)
	subscribe := new(SubscribeBcp)
	filelist, err := subscribe.WriteSubscribeBcp() //写bcp文件
	if err != nil {
		flog.Errorf("写入收藏bcp文件失败：%v \n", err)
		return
	}
	idx := new(index.Index)
	idx.BuildSubscribeIdx(filelist) //写索引文件

	bcpzip(filedir, model.SubscribeCode) //加密打包zip
}

func ZipUcar() {

	filedir := model.Basepath + model.UcarDir
	clean(filedir)
	ucar := new(UcarBcp)
	filelist, err := ucar.WriteUcarBcp() //写bcp文件
	if err != nil {
		flog.Errorf("写入车源发布bcp文件失败：%v \n", err)
		return
	}
	idx := new(index.Index)
	idx.BuildUcarIdx(filelist) //写索引文件

	bcpzip(filedir, model.UcarCode) //加密打包zip
}

func ZipEvaluate() {

	filedir := model.Basepath + model.EvaluateDir
	clean(filedir)
	ucar := new(EvaluateBcp)
	filelist, err := ucar.WriteEvaluateBcp() //写bcp文件
	if err != nil {
		flog.Errorf("写入车辆评估bcp文件失败：%v \n", err)
		return
	}
	idx := new(index.Index)
	idx.BuildEvaluateIdx(filelist) //写索引文件

	bcpzip(filedir, model.EvaluateCode) //加密打包zip
}

// 贷款购车
func ZipLoanOrder() {

	filedir := model.Basepath + model.LoanOrderDir
	clean(filedir)
	loanorder := new(LoanOrderBcp)
	filelist, err := loanorder.WriteLoanOrderBcp() //写bcp文件
	if err != nil {
		flog.Errorf("写入车辆评估bcp文件失败：%v \n", err)
		return
	}
	idx := new(index.Index)
	idx.BuildLoanOrderIdx(filelist) //写索引文件

	bcpzip(filedir, model.LoanOrderCode) //加密打包zip
}

// 贷款购车
func ZipBooks() {

	filedir := model.Basepath + model.BookDir
	clean(filedir)
	loanorder := new(BookBcp)
	filelist, err := loanorder.WriteBookBcp() //写bcp文件
	if err != nil {
		flog.Errorf("写入预约记录bcp文件失败：%v \n", err)
		return
	}
	idx := new(index.Index)
	idx.BuildBookIdx(filelist) //写索引文件

	bcpzip(filedir, model.BookCode) //加密打包zip
}

// 商户信息
func ZipDealers() {

	filedir := model.Basepath + model.DealerDir
	clean(filedir)
	dealer := new(DealerBcp)
	filelist, err := dealer.WriteDealerBcp() //写bcp文件
	if err != nil {
		flog.Errorf("写入商户信息bcp文件失败：%v \n", err)
		return
	}
	idx := new(index.Index)
	idx.BuildDealerIdx(filelist) //写索引文件

	bcpzip(filedir, model.DealerCode) //加密打包zip
}

// 写入文件 并返回文件列表
func writeBcp(total int64, dir, code string, getFileContent func(int64, int64) string) (map[string]int64, error) {

	filelist := make(map[string]int64)
	var start, end int64
	var bcpname string
	var pagecnt int64 = 1
	if total > pagesize {
		if total%pagesize == 0 {
			pagecnt = total / pagesize
		} else {
			pagecnt = 1 + total/pagesize
		}
	}

	now := time.Now()

	for i := int64(1); i <= pagecnt; i++ {

		start = (i-1)*pagesize + 1
		end = i * pagesize

		bcpname = strconv.Itoa(model.AppType) + "-" + tool.HandTime(now) + "-" + fmt.Sprintf("%05d", i) + "-" + code + "-0.bcp"
		if i == pagecnt {
			filelist[bcpname] = total % pagesize
		} else {
			filelist[bcpname] = pagesize
		}

		content := getFileContent(start, end)

		tdir := model.Basepath + dir
		if _, err := os.Open(tdir); err != nil {
			os.MkdirAll(tdir, os.ModePerm)
		}
		fpath := tdir + bcpname

		fileptr, err := os.OpenFile(fpath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		defer fileptr.Close()
		if err != nil {
			flog.Errorf("创建文件失败：%v \n", fpath, err)
			continue
		}

		fileptr.WriteString(strings.TrimSuffix(content, "\n"))
	}

	return filelist, nil
}

func bcpzip(filedir, code string) {

	now := time.Now()
	files, err := ioutil.ReadDir(filedir)
	if err != nil {
		flog.Errorf("打开用户bcp临时文件失败：%v \n", err)
		return
	}

	//zip path
	var path string
	datepath := fmt.Sprintf("%4d%02d%02d", now.Year(), now.Month(), now.Day())
	timepath := fmt.Sprintf("%02d%02d", now.Hour(), now.Minute())
	zipname := strconv.Itoa(model.AppType) + "-" + tool.HandTime(now) + "-11-1-00001.zip"

	rmdir := model.Basepath + model.OutputDir + datepath + string(os.PathSeparator) + code
	clean(rmdir)

	fdir := model.OutputDir + datepath + string(os.PathSeparator) + code + string(os.PathSeparator) + timepath + string(os.PathSeparator)
	path = model.Basepath + fdir
	if _, err := os.Open(path); err != nil {
		os.MkdirAll(path, os.ModePerm)
	}

	fzip, _ := os.Create(path + zipname)
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()

	pwd, _ := model.Cfg.String("zippwd")
	for _, file := range files {
		w, err := zipw.Encrypt(file.Name(), pwd)
		if err != nil {
			flog.Errorf("加密压缩zip包失败：%s %v \n", file.Name(), err)
			continue
		}
		contents, err := ioutil.ReadFile(filedir + file.Name())
		if err != nil {
			flog.Errorf("写入zip包时读取文件失败：%s %v \n", filedir+file.Name(), err)
			continue
		}
		_, err = io.Copy(w, bytes.NewReader(contents))
		if err != nil {
			flog.Errorf("写入zip包时copy文件数据流失败：%s %v \n", filedir+file.Name(), err)
			continue
		}
	}
	zipw.Flush()

	//zipw.Close() //  非defer，如果在当前方法内上传，需要手动关闭压缩流之后才能上传，否则上传的文件错误

	////上传ftp
	//server, _ := model.Cfg.String("ftp.server")
	//port, _ := model.Cfg.String("ftp.port")
	//username, _ := model.Cfg.String("ftp.username")
	//password, _ := model.Cfg.String("ftp.password")
	////ftpOp.UploadFile(strings.Join([]string{server, port}, ":"), username, password, path+zipname, fdir, zipname)
	//ftpOp.UploadFile_1(strings.Join([]string{server, port}, ":"), username, password, path+zipname, fdir, zipname)
}

//清空 目录下文件 重新生成
func clean(tdir string) {
	os.RemoveAll(tdir)
}
