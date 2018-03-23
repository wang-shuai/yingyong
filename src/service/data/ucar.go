package data

import (
	"fmt"
	"../model"
)

// 注册商户
func CountUCarDealer() (int64, error) {
	var u model.DealerAccount

	// 应该是一样的
	//total, err := ucar_engine.SQL(`select u.* from Dealer_Vendor_user u
	//				join Dealer_Vendor_Account a on a.DVAId = u.DVAId
	//				where u.Status=1 and a.Status !=-1`).Count(&u)
	total, err := ucar_engine.Table("Dealer_Vendor_Account").Alias("a").
	Join("inner",[]string{"Dealer_Vendor_user","u"},"a.DVAId = u.DVAId").
	Where("u.Status=1 and a.Status !=-1").Count(&u)

	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetUCarDealers(start, end int64) ([]model.DealerAccount, error) {
	var entities []model.DealerAccount
	// 要排序的所有数据
	all := `select ROW_NUMBER() over(order by a.CreateTime desc) as row,
		a.FullName as NAME,'0' as SEXCODE,'' as CERTIFICATE_TYPE,'' as CERTIFICATE_CODE,u.Mobile as MOBILE,'1430015' as REG_ACCOUNT_TYPE,
		a.DVAId as ACCOUNT_ID,u.UserName as REG_ACCOUNT,a.ShortName as REGIS_NICKNAME,a.CreateTime as REGIS_TIME,u.LastLoginIp as IP_ADDRESS,
		'' as PORT,'' as MAC_ADDRESS,a.Address as POSTAL_ADDRESS,a.Tel400 as CONTACTOR_TEL,a.CreateTime as BIRTHDAY,a.Pidname as COMPANY,
		'' as SAFE_QUESTION,'' as SAFE_ANSWER,'01' as ACTIVITE_TYPE,'' as ACTIVITE_ACCOUNT,u.Password as PASSWORD,'' as IMEI,'' as IMSI,
		a.BaiduMap as BAIDUMAP,'' as LONGITUDE,'' as LATITUDE,a.Address as SITE_ADDRESS,'' as ORIGIN_PLACE,'' as OFTEN_ADDRESS,
		a.DVAId as SHOP_ID,isnull(a.SiteDisplayName,a.ShortName) as SHOP_NAME,'' as DATA_LAND
		 from Dealer_Vendor_user u
		 join Dealer_Vendor_Account a on a.DVAId = u.DVAId
		where u.Status=1 and a.Status !=-1`
	sql := `select t.* from (%s) as t where t.row between %d and %d`

	err := ucar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}

// 收藏
func CountCollection() (int64, error) {
	var u model.Collection
	total, err := ucar_engine.Table("Iucar_Collection").Where("Status=1 and type =1 and CreateTime > '2016-01-01'").Count(&u)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetCollections(start, end int64) ([]model.Collection, error) {
	var entities []model.Collection
	all := `select ROW_NUMBER() over(order by c.[CreateTime] desc) as row,
		'' as SRC_IP,'' as DST_IP,'' as SRC_PORT,'' as DST_PORT,'' as MAC,c.CreateTime as CAPTURE_TIME,'' as IMSI,
		'' as EQUIPMENT_ID,'' as HARDWARE_SIGNATURE,'' as LONGITUDE,'' as LATITUDE,'02' as TERMINAL_TYPE,
		'' as TERMINAL_MODEL,'' as TERMINAL_OS_TYPE,'' as SOFTWARE_NAME,'' as DATA_LAND,'1430015' as APPLICATION_TYPE,
		isnull(c.UserID,'') as USER_INTENRALID,'' as USER_ACCOUNT,c.InfoID as NEWS_ID,'' as LABEL,'' as LIKE_TYPE,
		c.CreateTime as ACTION_TIME,'' as FILE_MD5,'' as FILE_ID,c.LinkUrl as FILE_URL,'E6' as ACTION_TYPE,
		c.InfoID as GOODS_ID,c.InfoDesc as GOODS_NAME,'' as GOODS_COMMENT,b.DisplayPrice as GOODS_PRICE,'' as SHOP_ID,
		'' as SHOP_NAME,'' as COLLECTPOSITION_NAME,'' as COLLECTPOSITION_ADDRESS from Iucar_Collection  c
		left join ucarbasicinfo b on c.InfoID=b.UcarID
		where c.Status=1 and c.type =1 and c.CreateTime > '2016-01-01'`
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := ucar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}
