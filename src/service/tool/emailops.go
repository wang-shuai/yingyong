package tool

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"strings"
	"crypto/tls"
	"gitlab.dev.daikuan.com/platform/golang-services/push-cities-to-redis/flog"
)

var emailConf *emailStruct

type emailStruct struct{
	host string
	port int
	username string
	password string
	from string
	to string
}

func EmailInit(host string,port int,username,password,from,to string) {
	emailConf = &emailStruct{
		host:host,
		port:port,
		username:username,
		password:password,
		from : from,
		to:to}
}


func SendMail(subject,body string) {
	m := gomail.NewMessage()

	tos := []string{"wangshuai@yxqiche.com"}
	if len(emailConf.to) > 0 {
		tos = strings.Split(emailConf.to, ",")
	}

	m.SetHeader("From", emailConf.from)
	m.SetHeader("To", tos...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(emailConf.host, emailConf.port, emailConf.username, emailConf.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("邮件发送成功！")
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		flog.Fatalf("发送邮件配置信息异常：%v", err)
	}
}
