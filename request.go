package nhtsa

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var nhtsaBaseUrl = "https://vpic.nhtsa.dot.gov/api/vehicles"

func doRequest[T any](rUrl string) (*T, error) {
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

	var resJson T

	jsonErr := json.Unmarshal(resBody, &resJson)
	if jsonErr != nil {
		return nil, fmt.Errorf("error unmarshalling response: %s", jsonErr)
	}

	return &resJson, nil
}
