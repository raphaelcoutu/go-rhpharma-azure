package routes

import (
	"crypto/sha256"
	"log"
	"strconv"
	"time"

	"code.soquee.net/otp"
	"github.com/gofiber/fiber/v2"
)

var secret = "123456abcde"

func Token(offset int) int32 {
	totp := otp.NewOTP([]byte(secret), 6, sha256.New, otp.TOTP(30*time.Second, func() time.Time {
		return time.Now()
	}))

	return totp(offset, nil)
}

func ValidateToken(value string) bool {
	token, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	return int32(token) == Token(0) || int32(token) == Token(-30)
}

func GetToken(c *fiber.Ctx) error {
	totpStr := strconv.FormatInt(int64(Token(0)), 10)

	return c.SendString(totpStr)
}
