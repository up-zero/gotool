package netutil

import "testing"

func TestMail_SendMail(t *testing.T) {
	m := Mail{
		Username: "用户名",
		Password: "密码",
		Addr:     "smtp.exmail.qq.com:465",
		From:     "name@example.com",
	}
	err := m.SendMail([]string{"getcharzp@qq.com"}, "title", "content")
	if err != nil {
		t.Log(err)
	}
}
