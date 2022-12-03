package define

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

var MailPassword = os.Getenv("MAIL_PD")

// 验证码长度
var CodeLength = 6

// 验证码过期时间 (s)
var CodeExpire = 300
