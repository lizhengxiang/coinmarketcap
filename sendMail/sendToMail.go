package sendMail

import (
	"strings"
	"net/smtp"
	"fmt"
	"strconv"
	"math"
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SendMail(Body string)  {
	err := SendToMail(User, Password, Host, To, Subject, Body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}



func MailTemplate(diff float64,cointype int){
	pow10_n := math.Pow10(2)
	diffResult := math.Trunc((diff+0.5/pow10_n)*pow10_n) / pow10_n
	float32s2 := strconv.FormatFloat(diffResult, 'f', -1, 64)//float64
	fmt.Println(float32s2)
	var Body = `
		<html>
			<body>
				<h3>
					`+"coin type " + strconv.Itoa(cointype)+" price fluctuation "+float32s2+`
				</h3>
			</body>
		</html>
		`
	SendMail(Body)
}
