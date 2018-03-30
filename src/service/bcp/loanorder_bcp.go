package bcp

import (
	"../model"
	"fmt"
	"strings"
	"../data"
	"../tool"
)

type LoanOrderBcp struct {
}

// 写入文件 并返回文件列表
func (this *LoanOrderBcp) WriteLoanOrderBcp() (map[string]int64, error) {

	cnt, err := data.CountLoanOrder()
	if err != nil {
		fmt.Println("获取用户总条数错误：", err)
		return nil, err
	}
	return writeBcp(cnt, model.LoanOrderDir, model.LoanOrderCode, getLoanOrderFileContent)
}

func getLoanOrderFileContent(start, end int64) string {

	var LoanOrders []model.LoanOrder
	LoanOrders, err := data.GetLoanOrders(start, end)
	if err != nil {
		fmt.Println("获取全部用户异常：", err)
		return ``
	}
	var content string
	for _, LoanOrder := range LoanOrders {

		LoanOrder.CAPTURE_TIME = tool.HandTimeStr(LoanOrder.CAPTURE_TIME)
		LoanOrder.CARD_TIME = tool.HandTimeStr(LoanOrder.CARD_TIME)

		LoanOrder.SRC_IP = tool.HandIP(LoanOrder.SRC_IP)
		LoanOrder.DST_IP = tool.HandIP(LoanOrder.DST_IP)

		line := strings.Join([]string{LoanOrder.SRC_IP,
			LoanOrder.DST_IP,
			LoanOrder.SRC_PORT,
			LoanOrder.DST_PORT,
			LoanOrder.MAC,
			LoanOrder.CAPTURE_TIME,
			LoanOrder.IMSI,
			LoanOrder.EQUIPMENT_ID,
			LoanOrder.HARDWARE_SIGNATURE,
			LoanOrder.LONGITUDE,
			LoanOrder.LATITUDE,
			LoanOrder.TERMINAL_TYPE,
			LoanOrder.TERMINAL_MODEL,
			LoanOrder.TERMINAL_OS_TYPE,
			LoanOrder.SOFTWARE_NAME,
			LoanOrder.DATA_LAND,
			LoanOrder.APPLICATION_TYPE,
			LoanOrder.ACCOUNT_ID,
			LoanOrder.ACCOUNT,
			LoanOrder.BUY_CITY,
			LoanOrder.Name,
			LoanOrder.SFZH,
			LoanOrder.BANK_ACCOUNT_NUM,
			LoanOrder.RELATIONSHIP_MOBILEPHONE,
			LoanOrder.CAREER_STYLE,
			LoanOrder.DRIVING_LICENSE,
			LoanOrder.LOAN,
			LoanOrder.CAR_BRAND,
			LoanOrder.CAR_TYPE,
			LoanOrder.CARD_TIME,
			LoanOrder.MILEAGE,
			LoanOrder.VEHICLE_CONDITION,
			LoanOrder.LICENSE_PLATE_SITE,
			LoanOrder.USED_CAR_PRICE,
		}, "\t")

		content += line + "\n"
	}
	//fmt.Println(content)
	return content
}
