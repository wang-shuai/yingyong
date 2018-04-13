package bcp

import (
	"golang-services/jingyong/model"
	"fmt"
	"strings"
	"golang-services/jingyong/data"
	"golang-services/jingyong/tool"

	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
)

type BookBcp struct {
}

// 写入文件 并返回文件列表
func (this *BookBcp) WriteBookBcp() (map[string]int64, error) {

	cnt, err := data.CountBooks()
	if err != nil {
		flog.Errorf("获取预约记录总条数错误：%v \n", err)
		return nil, err
	}
	return writeBcp(cnt, model.BookDir, model.BookCode, getBookFileContent)
}

func getBookFileContent(start, end int64) string {

	var entities []model.Book
	entities, err := data.GetBooks(start, end)
	if err != nil {
		flog.Errorf("获取全部预约记录异常：%v \n", err)
		return ``
	}
	var content string

	cnt := len(entities)
	turn := cnt / 100
	if cnt%100 != 0 {
		turn += 1
	}

	datamap := make(map[string]model.Entity_Cacheucarbasicinfo_Full)
	keyfmt := "BitAuto.Taoche.CarSource_FullDetail_UCarId_%d"
	for i := 0; i < turn; i++ {
		rediskeys := make([]interface{}, 0)

		for s := 100 * i; s < 100*(i+1) && s < cnt; s++ {
			book := entities[s]
			rediskeys = append(rediskeys, fmt.Sprintf(keyfmt, book.AimUcarId))
		}

		//books := make([]model.Entity_Cacheucarbasicinfo_Full, len(rediskeys))

		conn := model.Pool.Get()

		replys, err := redis.Values(conn.Do("MGET", rediskeys...))
		if err != nil {
			flog.Errorf("获取多个redis值出错：%v \n", err)
			continue
		}

		var item model.Entity_Cacheucarbasicinfo_Full
		for i, r := range replys {
			if bt, ok := r.([]byte); ok {
				json.Unmarshal(bt, &item)
				fmt.Println(i, "条目：", item)
				//books = append(books,item)
				datamap[item.UcarID] = item
			}
		}

		flog.Errorf("缓存字典：%v \n",datamap)
	}

	for i := range entities {
		entity := entities[i]
		if cbase, exist := datamap[entity.AimUcarId]; exist {
			entity.SRC_IP = tool.HandIP(entity.SRC_IP)
			entity.DST_IP = tool.HandIP(entity.DST_IP)
			entity.CAPTURE_TIME = tool.HandTimeStr(entity.CAPTURE_TIME)
			entity.CARD_TIME = tool.HandTimeStr(entity.CARD_TIME)
			entity.USED_CAR_URL = fmt.Sprintf("http://www.taoche.com/buycar/p-%s.html", entity.AimUcarId)
			//entity.ACCOUNT = cbase.UserID
			entity.ACCOUNT_ID = cbase.RegUserId
			entity.BINDING_PHONE = cbase.Tell

			entity.SELLER_PHONE = cbase.Tell2
			entity.USED_CAR_NAME = cbase.Serial_ShowName
			entity.USED_CAR_PRICE = cbase.DisplayPrice
			entity.CAR_BRAND = cbase.MBrand_Name
			entity.CAR_TYPE = cbase.CarType

			//entity.SALE_CITY   //sql 已赋值
			entity.MILEAGE = cbase.DrivingMileage
			entity.VEHICLE_CONDITION = ``
			entity.LICENSE_PLATE_SITE = cbase.LicenseCityId
		}

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
			entity.SELLER_PHONE,
			entity.USED_CAR_NAME,
			entity.USED_CAR_PRICE,
			entity.USED_CAR_URL,
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
	//flog.Errorf(content)
	return content
}
