package netutil

import (
	"crypto/tls"
	"net/smtp"
	"strings"
)

type Mail struct {
	Username string // 用户名
	Password string // 密码
	Addr     string // smtp地址, 例如 mail.example.com:smtp
	From     string // 发送者
}

// SendMail 发送邮件
//
// receiveUsers: 接收者
// subject: 主题
// content: 正文
func (mail Mail) SendMail(receiveUsers []string, subject, content string) error {
	// 跳过证书认证
	tlsConn, err := tls.Dial("tcp", mail.Addr, &tls.Config{InsecureSkipVerify: true, ServerName: strings.Split(mail.Addr, ":")[0]})
	if err != nil {
		return err
	}

	// 连接到SMTP服务器
	conn, err := smtp.NewClient(tlsConn, strings.Split(mail.Addr, ":")[0])
	if err != nil {
		return err
	}
	defer conn.Close()

	// 认证
	auth := smtp.PlainAuth("", mail.Username, mail.Password, strings.Split(mail.Addr, ":")[0])
	if err = conn.Auth(auth); err != nil {
		return err
	}
	// 设置发件人和收件人
	if err = conn.Mail(mail.Username); err != nil {
		return err
	}
	for _, recipient := range receiveUsers {
		if err = conn.Rcpt(recipient); err != nil {
			return err
		}
	}
	// 邮件内容
	wc, err := conn.Data()
	if err != nil {
		return err
	}
	message := "From: " + mail.From + "\r\n" +
		"To: " + strings.Join(receiveUsers, ";") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		content
	_, err = wc.Write([]byte(message))
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}

	return conn.Quit()
}
