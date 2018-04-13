package bcp

import (
	"golang-services/jingyong/model"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
	"strings"
	"golang-services/jingyong/data"
	"golang-services/jingyong/tool"
)

type DealerBcp struct {
}

// 写入文件 并返回文件列表
func (this *DealerBcp) WriteDealerBcp() (map[string]int64, error) {

	cnt, err := data.CountDealers()
	if err != nil {
		flog.Errorf("获取商户总条数错误：%v \n", err)
		return nil, err
	}
	return writeBcp(cnt, model.DealerDir, model.DealerCode, getDealerFileContent)
}

func getDealerFileContent(start, end int64) string {

	var entities []model.Dealer
	entities, err := data.GetDealers(start, end)
	if err != nil {
		flog.Errorf("获取全部商户异常：%v \n", err)
		return ``
	}
	var content string
	for _, entity := range entities {
		//经销商类型   1-经纪人	2-4S店  	3-专业公司  	4-厂商  	5-其他 	6-集团
		switch entity.DEALER_TYPE {
		case "1":
			entity.DEALER_TYPE = "经纪人"
		case "2":
			entity.DEALER_TYPE = "4S店"
		case "3":
			entity.DEALER_TYPE = "专业公司"
		case "4":
			entity.DEALER_TYPE = "厂商"
		case "5":
			entity.DEALER_TYPE = "其他"
		case "6":
			entity.DEALER_TYPE = "集团"
		}

		entity.SRC_IP = tool.HandIP(entity.SRC_IP)
		entity.DST_IP = tool.HandIP(entity.DST_IP)
		entity.CAPTURE_TIME = tool.HandTimeStr(entity.CAPTURE_TIME)

		line := strings.Join([]string{entity.SRC_IP,
			entity.DST_IP,
			entity.SRC_PORT,
			entity.DST_PORT,
			entity.MAC,
			entity.CAPTURE_TIME,
			entity.IMSI,
			entity.EQUIPMENT_ID,
			entity.HARDWARE_SIGNATURE,
			entity.LONGITUDE,
			entity.LATITUDE,
			entity.TERMINAL_TYPE,
			entity.TERMINAL_MODEL,
			entity.TERMINAL_OS_TYPE,
			entity.SOFTWARE_NAME,
			entity.DATA_LAND,
			entity.APPLICATION_TYPE,
			entity.ACCOUNT_ID,
			entity.ACCOUNT,
			entity.PRINCIPAL_NAME,
		    strings.Replace(entity.BINDING_PHONE,"\n","$",-1),
			entity.COMPANY_NAME,
			entity.DEALER_TYPE,
			entity.DEALER_ADDRESS,
		}, "\t")

		content += line + "\n"
	}

	return content
}
