package bcp

import (
	"golang-services/jingyong/model"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
	"strings"
	"golang-services/jingyong/data"
	"golang-services/jingyong/tool"
)

type SubscribeBcp struct {
}

// 写入文件 并返回文件列表
func (this *SubscribeBcp) WriteSubscribeBcp() (map[string]int64, error) {

	cnt, err := data.CountSubscribe()
	if err != nil {
		flog.Errorf("获取订阅总条数错误：%v \n", err)
		return nil, err
	}
	flog.Errorf("获取订阅记录总条数：%v \n", cnt)
	return writeBcp(cnt, model.SubscribeDir, model.SubscribeCode, getSubscribeFileContent)
}

func getSubscribeFileContent(start, end int64) string {

	var subscribes []model.Subscribe
	subscribes, err := data.GetSubscribes(start, end)
	if err != nil {
		flog.Errorf("获取全部订阅信息异常：%v \n", err)
		return ``
	}
	var content string
	for _, subscribe := range subscribes {
		subscribe.UPDATE_TIME = tool.HandTimeStr(subscribe.UPDATE_TIME)
		subscribe.CAPTURE_TIME = tool.HandTimeStr(subscribe.CAPTURE_TIME)
		subscribe.SRC_IP = tool.HandIP(subscribe.SRC_IP)
		subscribe.DST_IP = tool.HandIP(subscribe.DST_IP)

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
			subscribe.UPDATE_TIME,
			subscribe.B_IDEN_ID,
			subscribe.SUB_NAME,
			subscribe.SUB_ID,
			subscribe.SUB_NUM,
			subscribe.FILE_SIZE,
			subscribe.MAINFILE}, "\t")

		content += line + "\n"
	}

	return content
}
