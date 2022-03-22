package utils

import (
	"bytes"
	"go-starterkit-project/config"
	"go-starterkit-project/domain/data_models"
	"html/template"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(data *data_models.Mail) {
	var mailFrom string = data.From

	t := template.New(data.TemplateHtml)

	var errTpl error
	t, errTpl = t.ParseFiles("templates/" + data.TemplateHtml)
	if errTpl != nil {
		log.Println(errTpl)
	}

	var tpl bytes.Buffer
	if errTpl := t.Execute(&tpl, data.BodyParam); errTpl != nil {
		log.Println(errTpl)
	}

	if mailFrom == "" {
		mailFrom = config.Config("MAIL_FROM")
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", mailFrom)
	mailer.SetHeader("To", data.To...)
	mailer.SetHeader("Subject", data.Subject)
	mailer.SetBody("text/html", tpl.String())

	if data.Attachment != "" {
		mailer.Attach(data.Attachment)
	}

	mailPort, _ := strconv.ParseUint(config.Config("MAIL_PORT"), 10, 32)

	dialer := gomail.NewDialer(
		config.Config("MAIL_HOST"),
		int(mailPort),
		config.Config("MAIL_USERNAME"),
		config.Config("MAIL_PASSWORD"),
	)

	go func() {
		if err := dialer.DialAndSend(mailer); err != nil {
			log.Print(err)
		}
	}()
}
