package nhtsa

import "fmt"

func DecodeVin(vin string, year string) (*NHTSADecodeResponse[NHTSADecodeResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevin/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinFlat(vin string, year string) (*NHTSADecodeResponse[NHTSADecodeFlatResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvalues/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeFlatResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinExtended(vin string, year string) (*NHTSADecodeResponse[NHTSADecodeResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvaluesextended/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinFlatExtended(vin string, year string) (*NHTSADecodeResponse[NHTSADecodeFlatResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvaluesextended/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeFlatResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeWmi(wmi string) (*NHTSADecodeResponse[NHTSADecodeWmiResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodewmi/%s?format=json", nhtsaBaseUrl, wmi)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeWmiResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding WMI: %s", err)
	}

	return res, nil
}
