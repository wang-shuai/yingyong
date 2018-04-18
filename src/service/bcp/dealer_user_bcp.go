package bcp

import (
	"golang-services/jingyong/model"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
	"strings"
	"golang-services/jingyong/data"
	"golang-services/jingyong/tool"
)

type DealerUserBcp struct {
}

// 写入文件 并返回文件列表
func (this *DealerUserBcp) WriteDealerUserBcp() (map[string]int64, error) {

	cnt, err := data.CountUCarDealerUser()
	if err != nil {
		flog.Errorf("获取商户用户总条数错误：%v \n", err)
		return nil, err
	}
	flog.Errorf("获取商户用户总条数：%v \n", cnt)
	return writeBcp(cnt, model.DealerUserDir, model.DealerUserCode, getDealerUserFileContent)
}

func getDealerUserFileContent(start, end int64)string {

	var entities []model.DealerUser
	entities, err := data.GetUCarDealerUsers(start, end)
	if err != nil {
		flog.Errorf("获取全部商户用户异常：%v \n", err)
		return ``
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
	//flog.Errorf(content)

	return content
}
