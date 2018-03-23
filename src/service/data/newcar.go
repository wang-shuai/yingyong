package data

import (
	"fmt"
	"../model"
)

func CountAllUserInfos() (int64, error) {
	var u model.User
	total, err := newcar_engine.Table("LoanUser").Where("IsDeleted=0").Count(&u)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetAllUserInfos(start, end int64) ([]model.User, error) {
	var entities []model.User
	// 要排序的所有数据
	all := `select ROW_NUMBER() over(order by [CreateTime] desc) as row,
	Name,isnull(Gender,'男') as SexCode,111 as Certificate_Type,'' as Certificate_Code,Telphone as Mobile,'1430015' as Reg_Account_Type,
	ID as Account_Id,Telphone as Reg_Account,'昵称' as Regis_NickName, CreateTime as Regis_Time, isnull(IP,'') as Ip_Address,''as Port,
	'' as Mac_Address,'' as Postal_Address,Telphone as Contactor_Tel,'' as Birthday,'' as Company, '' as Safe_Question, '' as Safe_Answer,
	'03' as Activite_Type, Telphone as Activite_Account,Password,''as IMEI,'' as IMSI,'' as Longitude,'' as Latitude,
	'' as Site_Address,'' as Origin_Place,'' as Often_Address, '' as Data_Land
	 from LoanUser where IsDeleted=0`
	sql := `select t.* from (%s) as t where t.row between %d and %d`
	err := newcar_engine.SQL(fmt.Sprintf(sql, all, start, end)).Find(&entities)
	return entities, err
}

