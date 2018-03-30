package data

import (
	"fmt"
	"../model"
)

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

