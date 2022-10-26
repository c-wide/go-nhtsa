package nhtsa

import (
	"errors"
	"fmt"
)

func GetWmiForManufacturer(mId string, vType string) (*BaseApiResponse[GetWmiResult], error) {
	if mId == "" && vType == "" {
		return nil, errors.New("you must provide a manufacturer id or vehicle type")
	}

	reqUrl := fmt.Sprintf("%s/GetWMIsForManufacturer/%s", nhtsaBaseUrl, mId)

	if vType == "" {
		reqUrl += "?format=json"
	} else {
		reqUrl += fmt.Sprintf("?vehicleType=%s&format=json", vType)
	}

	res, err := doRequest[BaseApiResponse[GetWmiResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting WMI: %s", err)
	}

	return res, nil
}
