package data

import (
	"fmt"
	"../model"
)

// 注册商户
func CountUCarDealerUser() (int64, error) {
	var u model.DealerUser

	// 应该是一样的
	//total, err := ucar_engine.SQL(`select u.* from Dealer_Vendor_user u
	//				join Dealer_Vendor_Account a on a.DVAId = u.DVAId
	//				where u.Status=1 and a.Status !=-1`).Count(&u)
	total, err := ucar_engine.Table("Dealer_Vendor_Account").Alias("a").
		Join("inner", []string{"Dealer_Vendor_user", "u"}, "a.DVAId = u.DVAId").
		Join("inner", []string{"City", "c"}, "a.CityId = c.city_Id").
		Join("inner", []string{"Province", "p"}, "c.pvc_Id = p.pvc_Id").
		Where(fmt.Sprintf(`u.Status=1 and a.Status !=-1 and a.CreateTime > '%s'`,startdate)).Count(&u)

	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetUCarDealerUsers(start, end int64) ([]model.DealerUser, error) {
	var entities []model.DealerUser
	// 要排序的所有数据
	all := fmt.Sprintf(`select ROW_NUMBER() over(order by a.CreateTime desc) as row,
		a.FullName as NAME,'0' as SEXCODE,'' as CERTIFICATE_TYPE,'' as CERTIFICATE_CODE,u.Mobile as MOBILE,'1430015' as REG_ACCOUNT_TYPE,
		a.DVAId as ACCOUNT_ID,u.UserName as REG_ACCOUNT,a.ShortName as REGIS_NICKNAME,a.CreateTime as REGIS_TIME,u.LastLoginIp as IP_ADDRESS,
		'' as PORT,'' as MAC_ADDRESS,a.Address as POSTAL_ADDRESS,a.Tel400 as CONTACTOR_TEL,a.CreateTime as BIRTHDAY,a.Pidname as COMPANY,
		'' as SAFE_QUESTION,'' as SAFE_ANSWER,'01' as ACTIVITE_TYPE,'' as ACTIVITE_ACCOUNT,u.Password as PASSWORD,'' as IMEI,'' as IMSI,
		a.BaiduMap as BAIDUMAP,'' as LONGITUDE,'' as LATITUDE,a.Address as SITE_ADDRESS,'' as ORIGIN_PLACE,'' as OFTEN_ADDRESS,
		a.DVAId as SHOP_ID,isnull(a.SiteDisplayName,a.ShortName) as SHOP_NAME,'' as DATA_LAND,
		c.city_Name as DVR_CITY, '中国' as DVR_COUNTRY,c.city_Postcode as DVR_DEVISIONCODE,a.Address as DVR_ADRESSDETAIL,
		a.PostCode as DVR_POSTCODE,a.FullName as DVR_NAME,'' as DVR_ID,u.Mobile as DVR_MOBILE,u.Tel DVR_TELEPHONE,p.pvc_Name as DVR_PROVINCE,
		'' as DVR_AREA,'' as DVR_TOWN,'' as DVR_STATUS,'' as DVR_TOWN_NAME,'' as DVR_TOWN_CODE,'' as DEFAULTADDRESS,
		a.CreateTime as DVR_CAPTURETIME,a.LastModifyTime as DVR_UPDATETIME
		 from Dealer_Vendor_user u
		 join Dealer_Vendor_Account a on a.DVAId = u.DVAId
		 join City c on a.CityId=c.city_Id
		 join Province p on c.pvc_Id = p.pvc_Id
		where u.Status=1 and a.Status !=-1 and a.CreateTime > '%s'`,startdate)
	sql := `select t.* from (%s) as t where t.row between %d and %d`

	err := ucar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}

// 收藏
func CountCollection() (int64, error) {
	var u model.Collection
	total, err := ucar_engine.Table("Iucar_Collection").
		Where(fmt.Sprintf(`Status=1 and type =1 and CreateTime > '%s'`, startdate)).Count(&u)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetCollections(start, end int64) ([]model.Collection, error) {
	var entities []model.Collection
	all := fmt.Sprintf(`select ROW_NUMBER() over(order by c.CreateTime desc) as row,
		'' as SRC_IP,'' as DST_IP,'' as SRC_PORT,'' as DST_PORT,'' as MAC,c.CreateTime as CAPTURE_TIME,'' as IMSI,
		'' as EQUIPMENT_ID,'' as HARDWARE_SIGNATURE,'' as LONGITUDE,'' as LATITUDE,'02' as TERMINAL_TYPE,
		'' as TERMINAL_MODEL,'' as TERMINAL_OS_TYPE,'淘车' as SOFTWARE_NAME,'' as DATA_LAND,'1430015' as APPLICATION_TYPE,
		isnull(c.UserID,'') as USER_INTENRALID,'' as USER_ACCOUNT,c.InfoID as NEWS_ID,'' as LABEL,'' as LIKE_TYPE,
		c.CreateTime as ACTION_TIME,'' as FILE_MD5,'' as FILE_ID,c.LinkUrl as FILE_URL,'E6' as ACTION_TYPE,
		c.InfoID as GOODS_ID,c.InfoDesc as GOODS_NAME,'' as GOODS_COMMENT,b.DisplayPrice as GOODS_PRICE,'' as SHOP_ID,
		'' as SHOP_NAME,'' as COLLECTPOSITION_NAME,'' as COLLECTPOSITION_ADDRESS
		from Iucar_Collection  c
		left join ucarbasicinfo b on c.InfoID=b.UcarID
		where c.Status=1 and c.type =1 and c.CreateTime > '%s'`, startdate)
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := ucar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}

// 关注
func CountSubscribe() (int64, error) {
	var u model.Subscribe
	total, err := ucar_engine.Table("M_Subscribe").
		Where(fmt.Sprintf(`Status=1 and phone is not null and CreateTime > '%s'`, startdate)).Count(&u)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetSubscribes(start, end int64) ([]model.Subscribe, error) {
	var entities []model.Subscribe
	all := fmt.Sprintf(`select ROW_NUMBER() over(order by CreateTime desc) as row,
		'' as SRC_IP,'' as DST_IP,'' as SRC_PORT,'' as DST_PORT,'' as MAC,CreateTime as CAPTURE_TIME,'' as IMSI,
		'' as EQUIPMENT_ID,'' as HARDWARE_SIGNATURE,'' as LONGITUDE,'' as LATITUDE,'02' as TERMINAL_TYPE,
		'' as TERMINAL_MODEL,'' as TERMINAL_OS_TYPE,'淘车' as SOFTWARE_NAME,'' as DATA_LAND,'1430015' as APPLICATION_TYPE,
			'' as IDENTIFICATION_TYPE,isnull(UserId,'') as A_IDEN_ID,Phone as A_IDEN_STRING,Phone as A_PHONE,
			26 as ACTION_TYPE,LastAlterTime as UPDATE_time,'' as B_IDEN_ID,ResultFullUrl as SUB_NAME,KeyId as SUB_ID,
		(select count(1) from M_Subscribe  where Status=1 and phone=m.Phone group by Phone ) as SUB_NUM,
		'' as FILE_SIZE,'' as MAINFILE
		from M_Subscribe as m  where Status=1 and phone is not null and CreateTime > '%s'`, startdate)
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := ucar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}

// 评估
func CountUcar() (int64, error) {
	var u model.UcarBaseInfo
	total, err := ucar_engine.SQL(fmt.Sprintf(`select count(1) from UcarBasicInfo b
		join v_allcarbasic as p on b.CarID=p.Car_Id
		join City c on  b.CarCityId =c.city_Id
		join Dealer_Vendor_Account as a on b.UserID=a.DVAId
		 where b.CarPublishTime > '%s'  and b.Destroy=0 and b.UcarStatus > 0`,startdate)).Count(&u)

	if err != nil {
		return 0, err
	}
	//fmt.Println(total)
	return total, nil
}

func GetUcars(start, end int64) ([]model.UcarBaseInfo, error) {
	var entities []model.UcarBaseInfo
	all := fmt.Sprintf(`select ROW_NUMBER() over(order by b.CreateTime desc) as row,
		'' as SRC_IP,'' as DST_IP,'' as SRC_PORT,'' as DST_PORT,'' as MAC,b.CreateTime as CAPTURE_TIME,'' as IMSI,
		'' as EQUIPMENT_ID,'' as HARDWARE_SIGNATURE,'' as LONGITUDE,'' as LATITUDE,'02' as TERMINAL_TYPE,
		'' as TERMINAL_MODEL,'' as TERMINAL_OS_TYPE,'淘车' as SOFTWARE_NAME,'' as DATA_LAND,'1430015' as APPLICATION_TYPE,
		a.DVAId as ACCOUNT_ID, CONVERT(varchar, a.EPId ) +'@tc'  as REG_ACCOUNT, isnull(a.Tel400,a.CompanyPhone) as BINDING_PHONE,
		p.cb_Name + p.cs_Name + p.Car_Name as USED_CAR_NAME, b.DisplayPrice as USED_CAR_PRICE, '' as USED_CAR_URL,
		b.UcarID as USED_CAR_ID, p.cb_Name CAR_BRAND, b.CarType as CAR_TYPE, b.BuyCarDate as CARD_TIME,
		c.city_Name as SALE_CITY, b.DrivingMileage as MILEAGE, '' as VEHICLE_CONDITION, '' as LICENSE_PLATE_SITE
		from UcarBasicInfo as b
		join v_allcarbasic as p on b.CarID=p.Car_Id
		join City c on b.CarCityId =c.city_Id
		join Dealer_Vendor_Account as a on b.UserID=a.DVAId
		where b.CarPublishTime > '%s'  and b.Destroy=0 and b.UcarStatus > 0 and a.Status !=-1 `,startdate)
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := ucar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}

// 评估
func CountEvaluate() (int64, error) {
	var u model.EvaluateUcar
	total, err := ucar_engine.Table("evaluateUcarBasicInfo").Alias("e").
		Join("inner", []string{"evaluateRecord", "r"}, "e.EvalCarID=r.EvalCarID").
		Where(fmt.Sprintf(`e.EcarStatus=1 and e.Active=1 and e.CreateTime > '%[1]s'  and r.evaluateTime > '%[1]s'`,startdate)).Count(&u)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetEvaluates(start, end int64) ([]model.EvaluateUcar, error) {
	var entities []model.EvaluateUcar
	all := fmt.Sprintf(`select ROW_NUMBER() over(order by e.CreateTime desc) as row,
		'' as SRC_IP,'' as DST_IP,'' as SRC_PORT,'' as DST_PORT,'' as MAC,e.CreateTime as CAPTURE_TIME,'' as IMSI,
		'' as EQUIPMENT_ID,'' as HARDWARE_SIGNATURE,'' as LONGITUDE,'' as LATITUDE,'02' as TERMINAL_TYPE,
		'' as TERMINAL_MODEL,'' as TERMINAL_OS_TYPE,'淘车' as SOFTWARE_NAME,'' as DATA_LAND,'1430015' as APPLICATION_TYPE,
		''as ACCOUNT_ID,'' as ACCOUNT,v.cb_Name as CAR_BRAND,'' as CAR_TYPE,e.BuyCarDate as CARD_TIME,
		c.city_Name as SALE_CITY,e.DrivingMileage as MILEAGE,'' as VEHICLE_CONDITION,'' as LICENSE_PLATE_SITE,
		r.evaluatePrice as USED_CAR_PRICE,'' as TRANSFER_NUMBER,'' as EXPECTED_SELLING_TIME
		from evaluateUcarBasicInfo e
		join v_allcarbasic v on e.CarID=v.Car_Id
		join City c on e.EcarLocation=c.city_Id
		join evaluateRecord r on e.EvalCarID=r.EvalCarID
		where e.EcarStatus=1 and e.Active=1 and e.CreateTime > '%[1]s'  and r.evaluateTime > '%[1]s' `,startdate)
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := ucar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}

// 评估
func CountDealers() (int64, error) {
	var u model.Dealer
	total, err := ucar_engine.Table("Dealer_Vendor_Account").
		Where(fmt.Sprintf(`EPId>0 and Status!=-1 and CreateTime>'%s'`,startdate)).Count(&u)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetDealers(start, end int64) ([]model.Dealer, error) {
	var entities []model.Dealer
	all := fmt.Sprintf(`select ROW_NUMBER() over(order by CreateTime desc) as row,
		'' as SRC_IP,'' as DST_IP,'' as SRC_PORT,'' as DST_PORT,'' as MAC, CreateTime as CAPTURE_TIME,'' as IMSI,
		'' as EQUIPMENT_ID,'' as HARDWARE_SIGNATURE,'' as LONGITUDE,'' as LATITUDE,'02' as TERMINAL_TYPE,
		'' as TERMINAL_MODEL,'' as TERMINAL_OS_TYPE,'淘车' as SOFTWARE_NAME,'' as DATA_LAND,'1430015' as APPLICATION_TYPE,
		dvaid as ACCOUNT_ID, CONVERT(varchar, EPId ) +'@tc' as  ACCOUNT, ShortName as	PRINCIPAL_NAME,
 		isnull(Tel400,CompanyPhone) as  BINDING_PHONE,  FullName as COMPANY_NAME,  type as DEALER_TYPE,
  		Address as DEALER_ADDRESS
		from Dealer_Vendor_Account where EPId>0 and Status!=-1 and CreateTime>'%s'`,startdate)
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := ucar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}
