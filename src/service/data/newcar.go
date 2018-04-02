package data

import (
	"fmt"
	"../model"
)

// 注册用户
func CountAllUserInfos() (int64, error) {
	var u model.User
	total, err := newcar_engine.Table("LoanUser").
	Alias("u").Join("left",[]string{"LoanUserProfile","p"},"u.id = p.loanuserid").
	Where("u.IsDeleted=0 and p.IsDeleted=0").Count(&u)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetAllUserInfos(start, end int64) ([]model.User, error) {
	var entities []model.User
	// 要排序的所有数据
	all := `select ROW_NUMBER() over(order by u.CreateTime desc) as row,
	u.Name,isnull(u.Gender,'男') as SexCode,111 as Certificate_Type,p.CertificateNumber as Certificate_Code,u.Telphone as Mobile,'1430015' as Reg_Account_Type,
	u.ID as Account_Id,u.Telphone as Reg_Account,p.NickName as Regis_NickName, u.CreateTime as Regis_Time, isnull(u.IP,'') as Ip_Address,''as Port,
	'' as Mac_Address,p.Address as Postal_Address,u.Telphone as Contactor_Tel,p.Birthday as Birthday,'' as Company, '' as Safe_Question, '' as Safe_Answer,
	'03' as Activite_Type, u.Telphone as Activite_Account,u.Password,''as IMEI,'' as IMSI,'' as Longitude,'' as Latitude,
	'' as Site_Address,'' as Origin_Place,'' as Often_Address, '' as Data_Land ,
	93 as ACCOUNT_ACTION_TYPE,p.HeadPortrait as USER_PHOTO,	'' as HARDWARESTRING_TYPE,	'' as HARDWARESTRING,    '' as JOIN_ACCOUNT_TYPE,
	'' as IDEN_TYPE,    '' as IDENTIFICATION_TYPE,	'' as IDENTIFICATION_ID
	 from LoanUser as u
	 left join LoanUserProfile p on u.id = p.loanuserid
	 where u.IsDeleted=0 and p.IsDeleted=0`
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := newcar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}

// 贷款购车
func CountLoanOrder() (int64, error) {
	var u model.LoanOrder
	total, err := newcar_engine.Table("LoanOrder").Alias("o").
		Join("inner",[]string{"LoanUserProfile","u"},"o.userid = u.LoanUserID and u.IsDeleted=0").
		Join("inner",[]string{"LoanOrder_Middle","m"},"o.OrderID=m.LoanOrderID and m.IsDeleted=0").
		Where(fmt.Sprintf(`o.IsDeleted=0  and o.CreateTime > '%s'`,startdate)).Count(&u)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetLoanOrders(start, end int64) ([]model.LoanOrder, error) {
	var entities []model.LoanOrder
	// 要排序的所有数据
	all := fmt.Sprintf( `select ROW_NUMBER() over(order by u.CreateTime desc) as row,
		o.IP as SRC_IP,'' as DST_IP,'' as SRC_PORT,'' as DST_PORT,'' as MAC,o.CreateTime as CAPTURE_TIME,'' as IMSI,
		'' as EQUIPMENT_ID,'' as HARDWARE_SIGNATURE,'' as LONGITUDE,'' as LATITUDE,'02' as TERMINAL_TYPE,
		'' as TERMINAL_MODEL,'' as TERMINAL_OS_TYPE,'淘车' as SOFTWARE_NAME,'' as DATA_LAND,'1430015' as APPLICATION_TYPE,
		o.UserID as ACCOUNT_ID, o.Telphone as ACCOUNT,r.Name as BUY_CITY, u.RealName as Name,u.CertificateNumber as SFZH,
		'' as BANK_ACCOUNT_NUM,'' as RELATIONSHIP_MOBILEPHONE,'' as CAREER_STYLE,o.DrivingLicenceNumber as DRIVING_LICENSE,
		m.LoanMoney as LOAN,v.bs_Name as CAR_BRAND,'' as CAR_TYPE,'' as CARD_TIME,'' as MILEAGE,'' as VEHICLE_CONDITION,
		 '' as LICENSE_PLATE_SITE,'' as USED_CAR_PRICE
		from loanorder o
		join LoanUserProfile u on o.userid = u.LoanUserID and u.IsDeleted=0
		join Region r on o.CityID = r.ID and r.Level=2
		join LoanOrder_Middle m on o.OrderID=m.LoanOrderID and m.IsDeleted=0
		join ViewLevelCar v on o.CarId = v.Car_Id
		where o.IsDeleted=0  and o.CreateTime > '%s'`,startdate)
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := newcar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}