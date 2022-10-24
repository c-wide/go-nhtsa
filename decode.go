package nhtsa

import (
	"encoding/json"
	"fmt"
)

func DecodeVin(vin string, year string) (*NHTSADecodeResponse, error) {
	res, err := doRequest(fmt.Sprintf("%s/decodevin/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year))
	if err != nil {
		return nil, fmt.Errorf("error decoding vin: %s", err)
	}

	var resJson NHTSADecodeResponse

	jsonErr := json.Unmarshal(*res, &resJson)
	if jsonErr != nil {
		return nil, fmt.Errorf("error unmarshalling response: %s", jsonErr)
	}

	return &resJson, nil
}
