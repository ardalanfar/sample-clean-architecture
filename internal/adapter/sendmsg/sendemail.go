package sendmsg

import (
	"Farashop/internal/config"
	"Farashop/internal/contract"
	"context"
	"fmt"
	"net/smtp"
	"strings"
)

type Interactor struct {
	Config  *config.SendEmail
	To      []string
	Subject string
}

func NewSendMassage(to []string, subject string) contract.SendMassage {
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
	}
	return Interactor{
		Config:  cfg.Email,
		To:      to,
		Subject: subject,
	}
}

func (mail Interactor) SendEmail(ctx context.Context, msg string) error {
	auth := smtp.PlainAuth("", mail.Config.Username, mail.Config.Password, mail.Config.SmtpHost)

	err := smtp.SendMail(mail.Config.SmtpHost+":"+mail.Config.SmtpPort, auth, mail.Config.From, mail.To, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func (mail Interactor) BuildMessage() string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Config.From)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	return msg
}
