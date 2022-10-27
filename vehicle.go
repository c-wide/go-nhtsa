package nhtsa

import "fmt"

func GetMakesForVehicleType(vType string) (*BaseApiResponse[VehicleTypeNameResult], error) {
	reqUrl := fmt.Sprintf("%s/GetMakesForVehicleType/%s?format=json", nhtsaBaseUrl, vType)

	res, err := doRequest[BaseApiResponse[VehicleTypeNameResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting vehicle makes: %s", err)
	}

	return res, nil
}

func GetVehicleTypesForMakeByName(vMake string) (*BaseApiResponse[VehicleTypeNameResult], error) {
	reqUrl := fmt.Sprintf("%s/GetVehicleTypesForMake/%s?format=json", nhtsaBaseUrl, vMake)

	res, err := doRequest[BaseApiResponse[VehicleTypeNameResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting vehicle types: %s", err)
	}

	return res, nil
}

func GetVehicleTypesForMakeById(vId string) (*BaseApiResponse[VehicleTypeIdResult], error) {
	reqUrl := fmt.Sprintf("%s/GetVehicleTypesForMakeId/%s?format=json", nhtsaBaseUrl, vId)

	res, err := doRequest[BaseApiResponse[VehicleTypeIdResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting vehicle types: %s", err)
	}

	return res, nil
}
