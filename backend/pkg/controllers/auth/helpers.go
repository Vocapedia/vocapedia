package auth

import (
	"crypto/rand"
	"log"
	"strconv"
)

func generateOTP() string {
	otp := make([]byte, 6)
	_, err := rand.Read(otp)
	if err != nil {
		log.Println("Hata:", err)
		return ""
	}
	return strconv.Itoa(int(otp[0]%10)) + strconv.Itoa(int(otp[1]%10)) + strconv.Itoa(int(otp[2]%10)) +
		strconv.Itoa(int(otp[3]%10)) + strconv.Itoa(int(otp[4]%10)) + strconv.Itoa(int(otp[5]%10))
}
