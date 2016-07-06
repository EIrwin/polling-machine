package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

const (
	SMTP_USERNAME = "Eric.Irwin1124@gmail.com"
	SMTP_PASSWORD = "UvHw2w6jhpJl"
	SMTP_SERVER   = "mail.smtp2go.com"
	SMTP_PORT     = 2525
)

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, from string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
		from:    from,
	}
}

func (r *Request) SendEmail() (bool, error) {

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		SMTP_USERNAME,
		SMTP_PASSWORD,
		SMTP_SERVER,
	)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := fmt.Sprintf("%v:%v", SMTP_SERVER, SMTP_PORT)

	if err := smtp.SendMail(addr, auth, r.from, r.to, msg); err != nil {
		return false, err
	}

	return true, nil
}

func (r *Request) ParseTemplate(templatePath string, data interface{}) error {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println(err)
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Println(err)
		return err
	}
	r.body = buf.String()
	return nil

}
