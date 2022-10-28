package nhtsa

import (
	"fmt"
	"net/url"
)

func GetVehicleVariableList() (*BaseApiResponse[GetVehicleVariableResult], error) {
	reqUrl := fmt.Sprintf("%s/GetVehicleVariableList?format=json", nhtsaBaseUrl)

	res, err := doRequest[BaseApiResponse[GetVehicleVariableResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting variable list: %s", err)
	}

	return res, nil
}

func GetVehicleVariableValuesList(sVal string) (*BaseApiResponse[GetVehicleVariableValuesResult], error) {
	reqUrl := fmt.Sprintf("%s/GetVehicleVariableValuesList/", nhtsaBaseUrl)

	sValE := &url.URL{Path: sVal}

	res, err := doRequest[BaseApiResponse[GetVehicleVariableValuesResult]](reqUrl + sValE.String() + "?format=json")
	if err != nil {
		return nil, fmt.Errorf("error getting variable list: %s", err)
	}

	return res, nil
}
