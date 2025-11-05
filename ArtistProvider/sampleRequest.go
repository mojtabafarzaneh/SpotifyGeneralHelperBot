package artistprovider

import (
	"fmt"
	"io"
	"net/http"
)

func SampleRequest() {
	baseUrl := "https://api.spotify.com/v1/artists/4Z8W4fKeB5YxbusRsdQVPb"
	accessToken := "BQBs0hCwfX7i2X3VKqvPDNLz7Nw6nW_CJHOpxq5hL4kSvniSrax5Ib9-d7ktq4gt5t3UU_8ZIkvwFVwH7TRlMfvRhlCvrHYLUqYUwALsxeVRPt9jPgQP4k8yRVo6Ul71_Yr94rG-EI4"

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		fmt.Errorf("create request: %w", err)
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Request has been failed", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}
