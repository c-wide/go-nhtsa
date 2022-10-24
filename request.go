package nhtsa

import (
	"fmt"
	"io"
	"net/http"
)

var nhtsaBaseUrl = "https://vpic.nhtsa.dot.gov/api/vehicles"

func doRequest(rUrl string) (*[]byte, error) {
	req, err := http.NewRequest(http.MethodGet, rUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %s", err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", res.StatusCode)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	return &resBody, nil
}
