package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type GoogleAuthToken struct {
	AccessToken string
	IdToken     string
}

func GetGoogleAuthToken(code string) (*GoogleAuthToken, error) {
	const rootURL = "https://oauth2.googleapis.com/token"
	if err := godotenv.Load(); err != nil {
		ErrorHandler(err)
	}
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", os.Getenv("GOOGLE_CLIENT_ID"))
	values.Add("client_secret", os.Getenv("GOOGLE_CLIENT_SECRET"))
	values.Add("redirect_uri", os.Getenv("GOOGLE_REDIRECT_URI"))
	query := values.Encode()
	fmt.Println("query", query)
	req, err := http.NewRequest("POST", rootURL, bytes.NewBufferString(query))
	ErrorHandler(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}
	res, err := client.Do(req)
	fmt.Println("res: ", res)
	ErrorHandler(err)
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("could not retrieve users %v", res))
	}
	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleAuthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleAuthTokenRes); err != nil {
		return nil, err
	}
	fmt.Println("GoogleAuthToken: ", GoogleAuthTokenRes)
	tokenBody := &GoogleAuthToken{
		AccessToken: GoogleAuthTokenRes["access_token"].(string),
		IdToken:     GoogleAuthTokenRes["id_token"].(string),
	}

	return tokenBody, nil
}
