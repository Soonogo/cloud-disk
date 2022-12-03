package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity string, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	// false
	// token := jwt.NewWithClaims(jwt.SigningMethodES256, uc)
	// yse
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	fmt.Println(token, "2")
	fmt.Println("token www", token)

	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err

	}

	fmt.Println(tokenString, "11")
	return tokenString, nil
}

func MailSendCode(mail, code string) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := email.NewEmail()
	e.From = "Jordan Wright <tttsongen@foxmail.com>"
	e.To = []string{mail}
	//抄送
	// e.Bcc = []string{"test_bcc@example.com"}
	// e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Send Mail Test Subject" + " " + "Code  is " + code + "!"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Code  is " + code + "!</h1>")
	password := os.Getenv("MAIL_PD")
	// e.Send("smtp.qq.com:465", smtp.PlainAuth("", "test@gmail.com", "kfosodzlnonibdja", "smtp.qq.com"))
	err = e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "tttsongen@foxmail.com", password, "smtp.qq.com"), &tls.Config{
		InsecureSkipVerify: true, ServerName: "smtp.qq.com",
	})
	if err != nil {
		return err

	}
	return nil
}
