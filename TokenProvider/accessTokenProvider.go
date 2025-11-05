package TokenProvider

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func AccessTokenProvider() {
	tokenUrl := "https://accounts.spotify.com/api/token"
	const grantType = "client_credentials"
	const clientID = "712630d2b9424750a0ea6c9af45a2165"
	const clientSecret = "2166d6c5eec547ddb79b069e2869ffd9"

	data := url.Values{}

	data.Set("grant_type", grantType)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
