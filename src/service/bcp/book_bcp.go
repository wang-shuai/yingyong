package bcp

import (
	"../model"
	"fmt"
	"strings"
	"../data"
	"../tool"

	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

type BookBcp struct {
}

// 写入文件 并返回文件列表
func (this *BookBcp) WriteBookBcp() (map[string]int64, error) {

	cnt, err := data.CountBooks()
	if err != nil {
		fmt.Println("获取车辆评估总条数错误：", err)
		return nil, err
	}
	return writeBcp(cnt, model.BookDir, model.BookCode, getBookFileContent)
}

func getBookFileContent(start, end int64) string {

	var entities []model.Book
	entities, err := data.GetBooks(start, end)
	if err != nil {
		fmt.Println("获取全部车辆评估异常：", err)
		return ``
	}
	var content string

	cnt := len(entities)
	turn := cnt / 100
	if cnt%100 != 0 {
		turn += 1
	}

	keyfmt := "BitAuto.Taoche.CarSource_FullDetail_UCarId_%d"
	for i := 0; i < turn; i++ {
		rediskeys := make([]interface{}, 0)

		//datamap := make(map[int]model.Entity_Cacheucarbasicinfo_Full)

		for s := 100 * i; s < 100*(i+1) && s < cnt; s++ {
			book := entities[s]
			rediskeys = append(rediskeys, fmt.Sprintf(keyfmt, book.AimUcarId))
		}
		books := make([]model.Entity_Cacheucarbasicinfo_Full, len(rediskeys))

		conn := model.Pool.Get()

		replys, err := redis.Values(conn.Do("MGET", rediskeys...))
		if err != nil {
			fmt.Println("获取多个redis值出错：", err)
			continue
		}

		jsondata, err := json.Marshal(replys)
		if err == nil && len(jsondata) > 0 {
			fmt.Println(jsondata)
		} else {
			continue
		}
		json.Unmarshal(jsondata, books)

		fmt.Println(books)
	}

	for _, entity := range entities {

		entity.SRC_IP = tool.HandIP(entity.SRC_IP)
		entity.DST_IP = tool.HandIP(entity.DST_IP)
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
			//entity.TRANSFER_NUMBER,
			//entity.EXPECTED_SELLING_TIME
		},
			"\t")

		content += line + "\n"
	}
	//fmt.Println(content)
	return content
}
