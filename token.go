package main

import (
	"crypto/sha256"
	"strconv"
	"time"

	"code.soquee.net/otp"
	"github.com/gofiber/fiber/v2"
)

var secret = "123456abcde"

func Token() int32 {
	totp := otp.NewOTP([]byte(secret), 8, sha256.New, otp.TOTP(30*time.Second, func() time.Time {
		return time.Now()
	}))

	return totp(0, nil)
}

func GetToken(c *fiber.Ctx) error {
	totpStr := strconv.FormatInt(int64(Token()), 10)

	return c.SendString(totpStr)
}
