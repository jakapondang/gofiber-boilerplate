package mailpack

import (
	"bytes"
	"errors"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
)

var (
	host     string
	port     int
	username string
	password string
)

func SetConfig(cfg config.MailConfig) {
	host = cfg.Host
	port = cfg.Port
	username = cfg.Username
	password = cfg.Password

}

func SendMail(to, subject string, data any, templatePath string) error {
	rootDir, err := os.Getwd()
	if err != nil {
		return errors.New("(SendMail) Failed to read root dir:" + err.Error())

	}
	// Load the HTML template
	if templatePath == "" {
		return errors.New("Template Path cannot be empty")
	}
	tmpl, err := os.ReadFile(rootDir + "/pkg/utils/mailpack/templates/" + templatePath)
	if err != nil {
		return errors.New("(SendMail) Failed to read template file:" + err.Error())
	}

	// Parse the template
	t, err := template.New("mailpack").Parse(string(tmpl))
	if err != nil {
		return errors.New("(SendMail) Failed to parse template:" + err.Error())
	}

	// Execute the template and generate the mailpack body
	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return errors.New("(SendMail) Failed to execute template:" + err.Error())
	}

	// Create and send the mailpack
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@jakapondan.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(m); err != nil {
		return errors.New("(SendMail) Failed to send mailpack:" + err.Error())
	}
	return nil
}
