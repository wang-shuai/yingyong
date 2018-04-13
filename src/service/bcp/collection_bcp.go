package bcp

import (
	"golang-services/jingyong/model"
	"strings"
	"golang-services/jingyong/data"
	"golang-services/jingyong/tool"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
)

type CollectionBcp struct {
}

// 写入文件 并返回文件列表
func (this *CollectionBcp) WriteCollectionBcp() (map[string]int64, error) {

	cnt, err := data.CountCollection()
	if err != nil {
		flog.Errorf("获取商户总条数错误：%v \n", err)
		return nil, err
	}
	return writeBcp(cnt, model.CollectionDir, model.CollectionCode, getCollectionFileContent)
}

func getCollectionFileContent(start, end int64) string {

	var collections []model.Collection
	collections, err := data.GetCollections(start, end)
	if err != nil {
		flog.Errorf("获取全部收藏信息异常：%v \n", err)
		return ``
	}
	var content string
	for _, collection := range collections {
		collection.ACTION_TIME = tool.HandTimeStr(collection.ACTION_TIME)
		collection.CAPTURE_TIME = tool.HandTimeStr(collection.CAPTURE_TIME)
		collection.SRC_IP = tool.HandIP(collection.SRC_IP)
		collection.DST_IP = tool.HandIP(collection.DST_IP)

		line := strings.Join([]string{collection.SRC_IP,
			collection.DST_IP,
			collection.SRC_PORT,
			collection.DST_PORT,
			collection.MAC,
			collection.CAPTURE_TIME,
			collection.IMSI,
			collection.EQUIPMENT_ID,
			collection.HARDWARE_SIGNATURE,
			collection.LONGITUDE,
			collection.LATITUDE,
			collection.TERMINAL_TYPE,
			collection.TERMINAL_MODEL,
			collection.TERMINAL_OS_TYPE,
			collection.SOFTWARE_NAME,
			collection.DATA_LAND,
			collection.APPLICATION_TYPE,
			collection.USER_INTENRALID,
			collection.USER_ACCOUNT,
			collection.NEWS_ID,
			collection.LABEL,
			collection.LIKE_TYPE,
			collection.ACTION_TIME,
			collection.FILE_MD5,
			collection.FILE_ID,
			collection.FILE_URL,
			collection.ACTION_TYPE,
			collection.GOODS_ID,
			collection.GOODS_NAME,
			collection.GOODS_COMMENT,
			collection.GOODS_PRICE,
			collection.SHOP_ID,
			collection.SHOP_NAME,
			collection.COLLECTPOSITION_NAME,
			collection.COLLECTPOSITION_ADDRESS}, "\t")

		content += line + "\n"
	}
	return content
}
