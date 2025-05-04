package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/akifkadioglu/vocapedia/pkg/cache"
	"github.com/akifkadioglu/vocapedia/pkg/controllers/auth"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/test"
)

func TestSendOTP(t *testing.T) {
	s := test.TestInit()

	body, err := test.ConvertBody(auth.LoginBody{
		Email: "test@vocapedia.space",
	})
	if err != nil {
		t.Fatalf("Failed to convert body: %v", err)
	}

	req, _ := http.NewRequest("POST", "/api/v1/public/auth/send-otp", body)

	response := test.ExecuteRequest(req, s)

	test.CheckResponseCode(t, http.StatusOK, response.Code)
}

var verifyBody *bytes.Buffer

func TestVerifyOTP(t *testing.T) {
	ctx := context.Background()
	s := test.TestInit()
	rdb := cache.Redis()
	otp, err := rdb.Get(ctx, "otp:test@vocapedia.space").Result()

	if err != nil {
		log.Fatalln(err)
	}
	body, err := test.ConvertBody(auth.OtpBody{
		Email: "test@vocapedia.space",
		OTP:   otp,
		Device: entities.Device{
			Brand:    "test",
			Model:    "test",
			Os:       "test",
			Browser:  "test",
			Platform: "test",
			Language: "test",
		},
	})
	if err != nil {
		t.Fatalf("Failed to convert body: %v", err)
	}
	req, _ := http.NewRequest("POST", "/api/v1/public/auth/verify-otp", body)
	response := test.ExecuteRequest(req, s)

	test.CheckResponseCode(t, http.StatusOK, response.Code)

	verifyBody = response.Body
}

func TestLogout(t *testing.T) {
	s := test.TestInit()
	TestSendOTP(t)
	TestVerifyOTP(t)
	type VerifyOtpResponse struct {
		Token string `json:"token"`
	}
	var verifyOtpResponse VerifyOtpResponse

	err := test.ConvertBodyToStruct(verifyBody, &verifyOtpResponse)
	if err != nil {
		t.Fatalf("Failed to convert body: %v", err)
	}

	req, _ := http.NewRequest("DELETE", "/api/v1/auth/logout", nil)
	req.Header.Add("Authorization", "Bearer "+verifyOtpResponse.Token)
	response := test.ExecuteRequest(req, s)

	test.CheckResponseCode(t, http.StatusOK, response.Code)
}
