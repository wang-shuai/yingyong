package bcp

import (
	"../model"
	"fmt"
	"strings"
	"os"
	"../data"
)

type DealerBcp struct {
}

// 写入文件 并返回文件列表
func (this *DealerBcp) WriteDealerBcp() (map[string]int64, error) {

	cnt, err := data.CountUCarDealer()
	if err != nil {
		fmt.Println("获取商户总条数错误：", err)
		return nil, err
	}
	return writeBcp(cnt, model.DealerDir, model.DealerCode, writeDealerInfoToFile)
}

func writeDealerInfoToFile(start, end int64, bcpname string) {

	var users []model.DealerAccount
	users, err := data.GetUCarDealers(start, end)
	if err != nil {
		fmt.Println("获取全部商户异常：", err)
		return
	}
	var content string
	for _, user := range users {
		if len(user.BAIDUMAP) > 0 {
			ll := strings.Split(user.BAIDUMAP, ",")
			if len(ll) == 2 {
				user.LONGITUDE, user.LATITUDE = ll[0], ll[1]
			}
		}

		line := strings.Join([]string{user.NAME, user.SEXCODE, user.CERTIFICATE_TYPE, user.CERTIFICATE_CODE, user.MOBILE, user.REG_ACCOUNT_TYPE, user.ACCOUNT_ID,
			user.REG_ACCOUNT, user.REGIS_NICKNAME, user.REGIS_TIME, user.IP_ADDRESS, user.PORT, user.MAC_ADDRESS, user.POSTAL_ADDRESS,
			user.CONTACTOR_TEL, user.BIRTHDAY, user.COMPANY, user.SAFE_QUESTION, user.SAFE_ANSWER, user.ACTIVITE_TYPE,
			user.ACTIVITE_ACCOUNT, user.PASSWORD, user.IMEI, user.IMSI, user.LONGITUDE, user.LATITUDE, user.SITE_ADDRESS,
			user.ORIGIN_PLACE, user.OFTEN_ADDRESS, user.SHOP_ID, user.SHOP_NAME, user.DATA_LAND}, "\t")

		content += line + "\n"
	}
	//fmt.Println(content)

	dir := model.Basepath + model.DealerDir
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
