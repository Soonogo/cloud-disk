package test

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/jordan-wright/email"
)

func TestSendMail(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := email.NewEmail()
	e.From = "Jordan Wright <tttsongen@foxmail.com>"
	e.To = []string{"tttsongen@qq.com"}
	//抄送
	// e.Bcc = []string{"test_bcc@example.com"}
	// e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Send Mail Test Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	password := os.Getenv("MAIL_PD")
	// e.Send("smtp.qq.com:465", smtp.PlainAuth("", "test@gmail.com", "kfosodzlnonibdja", "smtp.qq.com"))
	err = e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "tttsongen@foxmail.com", password, "smtp.qq.com"), &tls.Config{
		InsecureSkipVerify: true, ServerName: "smtp.qq.com",
	})
	if err != nil {
		t.Fatal(err)

	}
}
