package utils

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/muhadyan/financial-planner/config"
	"github.com/muhadyan/financial-planner/model"
	"gopkg.in/gomail.v2"
)

func SendMail(templatePath string, data model.SendMail, subject string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(&body, data)

	m := gomail.NewMessage()
	m.SetHeader("From", config.GetConfig().SendFromAddress)
	m.SetHeader("To", data.SendTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, config.GetConfig().SendFromAddress, config.GetConfig().MailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
