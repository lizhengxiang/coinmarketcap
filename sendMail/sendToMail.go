package sendMail

import (
	"strings"
	"net/smtp"
	"fmt"
	"strconv"
	"math"
	"coinmarketcap/model"
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

func SendMail(Body,Subject string)  {
	fmt.Println(Subject)
	err := SendToMail(User, Password, Host, To, Subject, Body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}



func MailTemplate(diff float64,cointype int, Profit float64){
	pow10_n := math.Pow10(2)
	diffResult := math.Trunc((diff+0.5/pow10_n)*pow10_n) / pow10_n
	float32s2 := strconv.FormatFloat(diffResult, 'f', -1, 64)//float64

	profitResult := math.Trunc((Profit+0.5/pow10_n)*pow10_n) / pow10_n
	profits2 := strconv.FormatFloat(profitResult, 'f', -1, 64)//float64
	result := model.GetCurrencyInfoByCointype(cointype)
	var Body = `
		<html>
			<body>
				<h3>
					`+"coin type " + result.Name +" price fluctuation "+float32s2+" current Profit "+profits2+`
				</h3>
			</body>
		</html>
		`
	Subject := result.Name +" price fluctuation"
	SendMail(Body, Subject)
}
