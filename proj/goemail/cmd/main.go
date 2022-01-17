package main

import (
	"fmt"
	"log"

	"github.com/pescox/learning-go/proj/goemail"
)

func main() {
	emailHost := "smtp-mail.outlook.com"
	var emailPort int32 = 587
	emailUser := "med1tator@outlook.com"
	emailPassword := "Ding@2021"
	client, err := goemail.InitEmailClient(emailHost, emailPort, emailUser, emailPassword)
	if err != nil {
		log.Fatalln(err)
	}

	emailModel := goemail.EmailModel{
		Subject: "Golang Test Email",
		From:    "med1tator@outlook.com",
		To:      []string{"pescoding@outlook.com"},
		Cc:      []string{"dingpeng24001@talkweb.com.cn"},
		Body:    "This is an email from golang test client.",
	}
	err = client.SendEmail(emailModel)
	if err != nil {
		fmt.Println("send email err:", err)
		return
	}
	fmt.Println("please check mailbox")
}
