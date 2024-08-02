package mail

import (
	"bytes"
	"errors"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/ioutil"
)

type MailSender struct {
	Host     string
	Port     int
	Username string
	Password string
}

type MailData struct {
	Name           string
	ActivationLink string
}

func NewMailSender(config config.Config, templatePath string) *MailSender {
	return &MailSender{
		Host:     config.Mail.Host,
		Port:     config.Mail.Port,
		Username: config.Mail.Username,
		Password: config.Mail.Password,
	}
}

func (s *MailSender) SendMail(to, subject string, data MailData, templatePath string) error {
	// Load the HTML template
	if templatePath == "" {
		return errors.New("Template Path cannot be empty")
	}
	tmpl, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return errors.New("Failed to read template file:" + err.Error())
	}

	// Parse the template
	t, err := template.New("mail").Parse(string(tmpl))
	if err != nil {
		return errors.New("Failed to parse template:" + err.Error())
	}

	// Execute the template and generate the mail body
	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return errors.New("Failed to execute template:" + err.Error())
	}

	// Create and send the mail
	m := gomail.NewMessage()
	m.SetHeader("From", s.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)

	if err := d.DialAndSend(m); err != nil {
		return errors.New("Failed to send mail:" + err.Error())
	}
	return nil
}
