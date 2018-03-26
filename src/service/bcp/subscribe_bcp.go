package bcp

import (
	"../model"
	"fmt"
	"strings"
	"os"
	"../data"
)

type SubscribeBcp struct {
}

// 写入文件 并返回文件列表
func (this *SubscribeBcp) WriteSubscribeBcp() (map[string]int64, error) {

	cnt, err := data.CountSubscribe()
	if err != nil {
		fmt.Println("获取商户总条数错误：", err)
		return nil, err
	}
	return writeBcp(cnt, model.SubscribeDir, model.SubscribeCode, writeSubscribeInfoToFile)
}

func writeSubscribeInfoToFile(start, end int64, bcpname string) {

	var subscribes []model.Subscribe
	subscribes, err := data.GetSubscribes(start, end)
	if err != nil {
		fmt.Println("获取全部订阅信息异常：", err)
		return
	}
	var content string
	for _, subscribe := range subscribes {
		line := strings.Join([]string{subscribe.SRC_IP,
			subscribe.DST_IP,
			subscribe.SRC_PORT,
			subscribe.DST_PORT,
			subscribe.MAC,
			subscribe.CAPTURE_TIME,
			subscribe.IMSI,
			subscribe.EQUIPMENT_ID,
			subscribe.HARDWARE_SIGNATURE,
			subscribe.LONGITUDE,
			subscribe.LATITUDE,
			subscribe.TERMINAL_TYPE,
			subscribe.TERMINAL_MODEL,
			subscribe.TERMINAL_OS_TYPE,
			subscribe.SOFTWARE_NAME,
			subscribe.DATA_LAND,
			subscribe.APPLICATION_TYPE,
			subscribe.IDENTIFICATION_TYPE,
			subscribe.A_IDEN_ID,
			subscribe.A_IDEN_STRING,
			subscribe.A_PHONE,
			subscribe.ACTION_TYPE,
			subscribe.UPDATE_time,
			subscribe.B_IDEN_ID,
			subscribe.SUB_NAME,
			subscribe.SUB_ID,
			subscribe.SUB_NUM,
			subscribe.FILE_SIZE,
			subscribe.MAINFILE}, "\t")

		content += line + "\n"
	}

	dir := model.Basepath + model.SubscribeDir
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
