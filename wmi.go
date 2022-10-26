package nhtsa

import (
	"errors"
	"fmt"
	"net/url"
)

func GetWmiForManufacturer(mId string, vType string) (*BaseApiResponse[GetWmiResult], error) {
	if mId == "" && vType == "" {
		return nil, errors.New("you must provide a manufacturer id or vehicle type")
	}

	reqUrl := fmt.Sprintf("%s/GetWMIsForManufacturer/%s?", nhtsaBaseUrl, mId)
	queryParams := url.Values{}

	if vType != "" {
		queryParams.Add("vehicleType", vType)
	}

	queryParams.Add("format", "json")

	res, err := doRequest[BaseApiResponse[GetWmiResult]](reqUrl + queryParams.Encode())
	if err != nil {
		return nil, fmt.Errorf("error getting WMI: %s", err)
	}

	return res, nil
}
