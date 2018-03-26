package bcp

import (
	"../model"
	"fmt"
	"strings"
	"os"
	"../data"
	"time"
	"strconv"
)

type CollectionBcp struct {
}

// 写入文件 并返回文件列表
func (this *CollectionBcp) WriteCollectionBcp() (map[string]int64, error) {

	cnt, err := data.CountCollection()
	if err != nil {
		fmt.Println("获取商户总条数错误：", err)
		return nil, err
	}
	return writeBcp(cnt, model.CollectionDir, model.CollectionCode, writeCollectionInfoToFile)
}

func writeCollectionInfoToFile(start, end int64, bcpname string) {

	var collections []model.Collection
	collections, err := data.GetCollections(start, end)
	if err != nil {
		fmt.Println("获取全部收藏信息异常：", err)
		return
	}
	var content string
	for _, collection := range collections {
		if len(collection.ACTION_TIME) > 0 {
			the_time, err := time.ParseInLocation("2006-01-02 15:04:05", collection.ACTION_TIME, time.Local)
			if err == nil {
				collection.ACTION_TIME = strconv.FormatInt(the_time.Unix(),10)
			}else{
				collection.ACTION_TIME = strconv.FormatInt(time.Now().Unix(),10)
			}
		}

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

	dir := model.Basepath + model.CollectionDir
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
