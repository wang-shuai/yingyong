package bcp

import (
	"../model"
	"fmt"
	"strings"
	"os"
	"../data"
	"../tool"
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

	var entities []model.DealerAccount
	entities, err := data.GetUCarDealers(start, end)
	if err != nil {
		fmt.Println("获取全部商户异常：", err)
		return
	}
	var content string
	for _, entity := range entities {
		if len(entity.BAIDUMAP) > 0 {
			ll := strings.Split(entity.BAIDUMAP, ",")
			if len(ll) == 2 {
				entity.LONGITUDE, entity.LATITUDE = ll[0], ll[1]
			}
		}
		entity.REGIS_TIME = tool.HandTimeStr(entity.REGIS_TIME)
		entity.IP_ADDRESS = tool.HandIP(entity.IP_ADDRESS)
		entity.DVR_CAPTURETIME = tool.HandTimeStr(entity.DVR_CAPTURETIME)
		entity.DVR_UPDATETIME = tool.HandTimeStr(entity.DVR_UPDATETIME)

		line := strings.Join([]string{entity.NAME, entity.SEXCODE, entity.CERTIFICATE_TYPE, entity.CERTIFICATE_CODE, entity.MOBILE, entity.REG_ACCOUNT_TYPE, entity.ACCOUNT_ID,
			entity.REG_ACCOUNT, entity.REGIS_NICKNAME, entity.REGIS_TIME, entity.IP_ADDRESS, entity.PORT, entity.MAC_ADDRESS, entity.POSTAL_ADDRESS,
			entity.CONTACTOR_TEL, entity.BIRTHDAY, entity.COMPANY, entity.SAFE_QUESTION, entity.SAFE_ANSWER, entity.ACTIVITE_TYPE,
			entity.ACTIVITE_ACCOUNT, entity.PASSWORD, entity.IMEI, entity.IMSI, entity.LONGITUDE, entity.LATITUDE, entity.SITE_ADDRESS,
			entity.ORIGIN_PLACE, entity.OFTEN_ADDRESS, entity.SHOP_ID, entity.SHOP_NAME, entity.DATA_LAND, entity.DVR_CITY,
			entity.DVR_COUNTRY, entity.DVR_DEVISIONCODE, entity.DVR_ADRESSDETAIL, entity.DVR_POSTCODE, entity.DVR_NAME,
			entity.DVR_ID, entity.DVR_MOBILE, entity.DVR_TELEPHONE, entity.DVR_PROVINCE, entity.DVR_AREA, entity.DVR_TOWN,
			entity.DVR_STATUS, entity.DVR_TOWN_NAME, entity.DVR_TOWN_CODE, entity.DEFAULTADDRESS, entity.DVR_CAPTURETIME,
			entity.DVR_UPDATETIME}, "\t")

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
