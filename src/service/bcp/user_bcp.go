package bcp

import (
	"golang-services/jingyong/model"
	"strings"
	"golang-services/jingyong/data"
	"golang-services/jingyong/tool"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
)

type UserBcp struct {
}

// 写入文件 并返回文件列表
func (this *UserBcp) WriteUserBcp() (map[string]int64, error) {

	cnt, err := data.CountAllUserInfos()
	if err != nil {
		flog.Errorf("获取用户总条数错误：%v \n", err)
		return nil, err
	}
	return writeBcp(cnt, model.UserDir, model.UserCode, getUserFileContent)
}

func getUserFileContent(start, end int64) string {

	var users []model.User
	users, err := data.GetAllUserInfos(start, end)
	if err != nil {
		flog.Errorf("获取全部用户异常：%v \n", err)
		return ``
	}
	var content string
	for _, user := range users {
		//line :=  user.Name + "\t" + user.SexCode + "\t" + user.Certificate_Type + "\t" + user.Certificate_Code + "\t" + user.Mobile + "\t" + user.Reg_Account_Type + "\t" + user.Account_Id + "\t" + user.Reg_Account + "\t" + user.Regis_NickName + "\t" + user.Regis_Time + "\t" + user.Ip_Address + "\t" + user.Port + "\t" + user.Mac_Address + "\t" + user.Postal_ddress + "\t" + user.Contactor_Tel + "\t" + user.Birthday + "\t" + user.Company + "\t" + user.Safe_Question + "\t" + user.Safe_Answer + "\t" + user.Activite_Type + "\t" + user.Activite_Account + "\t" + user.Password + "\t" + user.IMEI + "\t" + user.IMSI + "\t" + user.Longitude + "\t" + user.Latitude + "\t" + user.Site_Address + "\t" + user.Origin_Place + "\t" + user.Often_Address + "\t" + user.Data_Land
		user.Regis_Time = tool.HandTimeStr(user.Regis_Time)
		user.Ip_Address = tool.HandIP(user.Ip_Address)

		line := strings.Join([]string{user.Name, user.SexCode, user.Certificate_Type, user.Certificate_Code, user.Mobile, user.Reg_Account_Type, user.Account_Id,
			user.Reg_Account, user.Regis_NickName, user.Regis_Time, user.Ip_Address, user.Port, user.Mac_Address, user.Postal_Address,
			user.Contactor_Tel, user.Birthday, user.Company, user.Safe_Question, user.Safe_Answer, user.Activite_Type,
			user.Activite_Account, user.Password, user.IMEI, user.IMSI, user.Longitude, user.Latitude, user.Site_Address,
			user.Origin_Place, user.Often_Address, user.Data_Land, user.ACCOUNT_ACTION_TYPE, user.USER_PHOTO,
			user.HARDWARESTRING_TYPE, user.HARDWARESTRING, user.JOIN_ACCOUNT_TYPE, user.IDEN_TYPE, user.IDENTIFICATION_TYPE,
			user.IDENTIFICATION_ID}, "\t")

		content += line + "\n"
	}
	//flog.Errorf(content)

	return content
}
