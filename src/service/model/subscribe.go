package model

// 订阅
type Subscribe struct {
	Base

	IDENTIFICATION_TYPE string //合作账号类型
	A_IDEN_ID           string //用户ID
	A_IDEN_STRING       string //用户账号
	A_PHONE             string //手机号
	ACTION_TYPE         string //操作类型（关注/订阅/取消关注/取消订阅）
	UPDATE_TIME         string //操作时间
	B_IDEN_ID           string //被关注/订阅用户ID
	SUB_NAME            string //关注/订阅名称（文件/频道）
	SUB_ID              string //关注/订阅（文件/频道）ID
	SUB_NUM             string //关注/订阅数量
	FILE_SIZE           string //关注/订阅文件大小
	MAINFILE            string //关注/订阅文件实体
}
