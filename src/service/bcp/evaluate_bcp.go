package bcp

import (
	"golang-services/jingyong/model"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
	"strings"
	"golang-services/jingyong/data"
	"golang-services/jingyong/tool"
)

type EvaluateBcp struct {
}

// 写入文件 并返回文件列表
func (this *EvaluateBcp) WriteEvaluateBcp() (map[string]int64, error) {

	cnt, err := data.CountEvaluate()
	if err != nil {
		flog.Errorf("获取车辆评估总条数错误：%v \n", err)
		return nil, err
	}
	return writeBcp(cnt, model.EvaluateDir, model.EvaluateCode, getEvaluateFileContent)
}

func getEvaluateFileContent(start, end int64) string {

	var entities []model.EvaluateUcar
	entities, err := data.GetEvaluates(start, end)
	if err != nil {
		flog.Errorf("获取全部车辆评估异常：%v \n", err)
		return ``
	}
	var content string
	for _, entity := range entities {

		entity.SRC_IP = tool.HandIP(entity.SRC_IP)
		entity.DST_IP = tool.HandIP(entity.DST_IP)
		entity.EXPECTED_SELLING_TIME = tool.HandTimeStr(entity.EXPECTED_SELLING_TIME)
		entity.CAPTURE_TIME = tool.HandTimeStr(entity.CAPTURE_TIME)
		entity.CARD_TIME = tool.HandTimeStr(entity.CARD_TIME)

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
			entity.CAR_BRAND,
			entity.CAR_TYPE,
			entity.CARD_TIME,
			entity.SALE_CITY,
			entity.MILEAGE,
			entity.VEHICLE_CONDITION,
			entity.LICENSE_PLATE_SITE,
			entity.USED_CAR_PRICE,
			entity.TRANSFER_NUMBER,
			entity.EXPECTED_SELLING_TIME}, "\t")

		content += line + "\n"
	}
	//flog.Errorf(content)
	return content
}
