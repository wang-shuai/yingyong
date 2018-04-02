package data

import (
	"fmt"
	"../model"
)

// 预约记录
func CountBooks() (int64, error) {
	var u model.Book
	total, err := uleads_engine.Table("Leads_LeadsInfo").Alias("l").
		Join("inner",[]string{"AreaCity","c"},"l.CityID=c.CityId").
		Where(fmt.Sprintf(`l.CreateTime > '%s' and l.AimUcarId > 0
		and l.Category in (1010,1011,1012,1014,2014,5040,5010,5020,5030,3010,5080,5060,8010,8020,8030,8040)`,startdate)).Count(&u)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetBooks(start, end int64) ([]model.Book, error) {
	var entities []model.Book
	all := fmt.Sprintf(`select ROW_NUMBER() over(order by l.CreateTime desc) as row,
		'' as SRC_IP,'' as DST_IP,'' as SRC_PORT,'' as DST_PORT,'' as MAC,l.CreateTime as CAPTURE_TIME,'' as IMSI,
		'' as EQUIPMENT_ID,'' as HARDWARE_SIGNATURE,'' as LONGITUDE,'' as LATITUDE,'02' as TERMINAL_TYPE,
		'' as TERMINAL_MODEL,'' as TERMINAL_OS_TYPE,'淘车' as SOFTWARE_NAME,'' as DATA_LAND,'1430015' as APPLICATION_TYPE,
		'' as ACCOUNT_ID,'' as ACCOUNT,l.Mobile as BINDING_PHONE,'' as SELLER_PHONE,'' as USED_CAR_NAME,
		'' as USED_CAR_PRICE,'' as USED_CAR_URL,'' as CAR_BRAND,'' as CAR_TYPE,'' as CARD_TIME,c.CityAreaId as SALE_CITY,
		'' as MILEAGE,'' as VEHICLE_CONDITION,'' as LICENSE_PLATE_SITE,AimUcarId
		from Leads_LeadsInfo as l
		join AreaCity as c on l.CityID=c.CityId
		where l.CreateTime > '%s' and l.AimUcarId >0
		and l.Category in (1010,1011,1012,1014,2014,5040,5010,5020,5030,3010,5080,5060,8010,8020,8030,8040) `,startdate)
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := uleads_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}