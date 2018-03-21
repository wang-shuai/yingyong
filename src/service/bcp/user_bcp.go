package bcp

import (
	"../model"
	"fmt"
	//"strings"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)
type UserBcp struct {
}

// 写入文件 并返回文件列表
func (this * UserBcp) WriteUserBcp() ([]string,error) {
	clean()

	cnt, err := CountAllUserInfos()
	if err != nil {
		fmt.Println("获取用户总条数错误：", err)
		return nil,err
	}
	//fmt.Println("注册用户总数量：", cnt)

	var filelist []string = make([]string, 0)

	var start, end int64
	var bcpname string
	var pagecnt int64 = 1
	if cnt > pagesize {
		if cnt%pagesize == 0 {
			pagecnt = cnt / pagesize
		} else {
			pagecnt = 1 + cnt/pagesize
		}
	}

	now := time.Now()
	timespan := now.Unix()
	var wg sync.WaitGroup
	for i := int64(1); i <= pagecnt; i++ {

		start = (i - 1) * pagesize
		end = start + pagesize

		bcpname = strconv.Itoa(model.AppType) + "-" + strconv.FormatInt(timespan, 10) + "-" + fmt.Sprintf("%05d", i) + "-" + model.UserCode + "-0.bcp"
		filelist = append(filelist, bcpname)

		wg.Add(1)
		go func(start, end int64, name string) {
			defer wg.Done() //wg.Add(-1)
			writeUserInfoToFile(start, end, name)
		}(start, end, bcpname)
	}
	wg.Wait()
	fmt.Println("所有的线程执行结束")

	return filelist,nil
}

func writeUserInfoToFile(start, end int64, bcpname string) {

	var users []model.User
	users, err := GetAllUserInfos(start, end)
	if err != nil {
		fmt.Println("获取全部用户异常：", err)
		return
	}
	var content string
	for _, user := range users {
		//line :=  user.Name + "\t" + user.SexCode + "\t" + user.Certificate_Type + "\t" + user.Certificate_Code + "\t" + user.Mobile + "\t" + user.Reg_Account_Type + "\t" + user.Account_Id + "\t" + user.Reg_Account + "\t" + user.Regis_NickName + "\t" + user.Regis_Time + "\t" + user.Ip_Address + "\t" + user.Port + "\t" + user.Mac_Address + "\t" + user.Postal_ddress + "\t" + user.Contactor_Tel + "\t" + user.Birthday + "\t" + user.Company + "\t" + user.Safe_Question + "\t" + user.Safe_Answer + "\t" + user.Activite_Type + "\t" + user.Activite_Account + "\t" + user.Password + "\t" + user.IMEI + "\t" + user.IMSI + "\t" + user.Longitude + "\t" + user.Latitude + "\t" + user.Site_Address + "\t" + user.Origin_Place + "\t" + user.Often_Address + "\t" + user.Data_Land

		line := strings.Join([]string{user.Name, user.SexCode, user.Certificate_Type, user.Certificate_Code, user.Mobile, user.Reg_Account_Type, user.Account_Id,
			user.Reg_Account, user.Regis_NickName, user.Regis_Time, user.Ip_Address, user.Port, user.Mac_Address, user.Postal_ddress,
			user.Contactor_Tel, user.Birthday, user.Company, user.Safe_Question, user.Safe_Answer, user.Activite_Type,
			user.Activite_Account, user.Password, user.IMEI, user.IMSI, user.Longitude, user.Latitude, user.Site_Address,
			user.Origin_Place, user.Often_Address, user.Data_Land}, "\t")

		content += line + "\n"
	}
	//fmt.Println(content)

	base, _ := os.Getwd()
	dir := base + model.UserDir
	if _, err := os.Open(dir); err != nil {
		os.MkdirAll(dir, os.ModePerm)
	}
	fpath := dir + bcpname

	fileptr, err := os.OpenFile(fpath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer fileptr.Close()
	if err != nil {
		fmt.Println("创建文件失败：", fpath, err)
		return
	}

	fileptr.WriteString(content)
}

//清空 目录下文件 重新生成
func clean(){
	base, _ := os.Getwd()
	dir := base + model.UserDir
	os.RemoveAll(dir)
}