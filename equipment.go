package nhtsa

import (
	"fmt"
	"net/url"
)

func GetEquipmentPlantCodes(year string, eType string, rType string) (*BaseApiResponse[GetEquipmentPlantCodesResult], error) {
	reqUrl := fmt.Sprintf("%s/GetEquipmentPlantCodes?", nhtsaBaseUrl)

	queryParams := url.Values{}

	if year != "" {
		queryParams.Add("year", year)
	}

	if eType != "" {
		queryParams.Add("equipmentType", eType)
	}

	if rType != "" {
		queryParams.Add("reportType", rType)
	}

	queryParams.Add("format", "json")

	res, err := doRequest[BaseApiResponse[GetEquipmentPlantCodesResult]](reqUrl + queryParams.Encode())
	if err != nil {
		return nil, fmt.Errorf("error getting equipment plant codes: %s", err)
	}

	return res, nil
}
