package mailer

import (
	"bytes"
	"embed"
	"html/template"
	"time"

	"github.com/go-mail/mail/v2"
)

//go:embed "templates"
var templateFS embed.FS

// Mailer struct, contains a mail.Dialer instance and the sender
// information for the email to be from
type Mailer struct {
	dialer *mail.Dialer
	sender string
}

// Returns a new Dialer instance which contains the dialer and sender information.
func New(host string, port int, username, password, sender string) Mailer {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 10 * time.Second

	return Mailer{
		dialer: dialer,
		sender: sender,
	}
}

// Takes the recipent email address as the first parameter, the name of the
// containing the templates, and any dynamic data for the templates as any parameter.
func (m Mailer) Send(recipent, templateFile string, data any) error {
	tmpl, err := template.New("email").ParseFS(templateFS,
		"templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return nil
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}

	// Creating the mail
	msg := mail.NewMessage()
	msg.SetHeader("To", recipent)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	for i := 1; i <= 3; i++ {
		err = m.dialer.DialAndSend(msg)
		if nil == err {
			return nil
		}

		time.Sleep(1300 * time.Millisecond)
	}

	return nil
}
