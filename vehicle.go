package nhtsa

import "fmt"

func GetMakesForVehicleType(vType string) (*BaseApiResponse[GetMakeVehicleTypeNameResult], error) {
	reqUrl := fmt.Sprintf("%s/GetMakesForVehicleType/%s?format=json", nhtsaBaseUrl, vType)

	res, err := doRequest[BaseApiResponse[GetMakeVehicleTypeNameResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting vehicle makes: %s", err)
	}

	return res, nil
}
