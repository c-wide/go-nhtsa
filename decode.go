package nhtsa

import "fmt"

func DecodeVin(vin string, year string) (*NHTSADecodeResponse, error) {
	res, err := doRequest[NHTSADecodeResponse](fmt.Sprintf("%s/decodevin/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year))
	if err != nil {
		return nil, fmt.Errorf("error decoding vin: %s", err)
	}

	return res, nil
}
