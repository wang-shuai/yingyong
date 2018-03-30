package bcp

import (
	"../model"
	"fmt"
	"strings"
	"../data"
	"../tool"
)

type UcarBcp struct {
}

// 写入文件 并返回文件列表
func (this *UcarBcp) WriteUcarBcp() (map[string]int64, error) {

	cnt, err := data.CountUcar()
	if err != nil {
		fmt.Println("获取车辆评估总条数错误：", err)
		return nil, err
	}
	return writeBcp(cnt, model.UcarDir, model.UcarCode, getUcarFileContent)
}

func getUcarFileContent(start, end int64) string {

	var entities []model.UcarBaseInfo
	entities, err := data.GetUcars(start, end)
	if err != nil {
		fmt.Println("获取全部车源异常：", err)
		return ``
	}
	var content string
	for _, entity := range entities {

		entity.SRC_IP = tool.HandIP(entity.SRC_IP)
		entity.DST_IP = tool.HandIP(entity.DST_IP)
		entity.CAPTURE_TIME = tool.HandTimeStr(entity.CAPTURE_TIME)
		entity.CARD_TIME = tool.HandTimeStr(entity.CARD_TIME)
		entity.USED_CAR_URL = fmt.Sprintf("http://www.taoche.com/buycar/p-%s.html",entity.USED_CAR_ID)

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
			entity.BINDING_PHONE,
			entity.USED_CAR_NAME,
			entity.USED_CAR_PRICE,
			entity.USED_CAR_URL,
			entity.USED_CAR_ID,
			entity.CAR_BRAND,
			entity.CAR_TYPE,
			entity.CARD_TIME,
			entity.SALE_CITY,
			entity.MILEAGE,
			entity.VEHICLE_CONDITION,
			entity.LICENSE_PLATE_SITE,
		}, "\t")

		content += line + "\n"
	}
	return content
}
