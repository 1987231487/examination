package util
import (
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func SentEmail(addr  string)(code string,err error){
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code = fmt.Sprintf("%06v", rnd.Int31n(1000000))

	m := gomail.NewMessage()
	m.SetHeader("From", "1987231487@qq.com") //发件人
	m.SetHeader("To", addr)       //收件人
	//m.SetAddressHeader("Cc", "test@126.com", "test")     //抄送人
	m.SetHeader("Subject", "获取您的验证码")      //邮件标题
	m.SetBody("text/html", "您的验证码为:"+code) //邮件内容
	//m.Attach("E:\\IMGP0814.JPG")       //邮件附件

	d := gomail.NewDialer("smtp.qq.com", 465, "1987231487@qq.com", "xxxxxxxxxxx")
	//邮件发送服务器信息,使用授权码而非密码
	if err := d.DialAndSend(m); err != nil {
		return "",err
	}
	return code,err
}
