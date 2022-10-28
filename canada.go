package nhtsa

import (
	"fmt"
	"net/url"
	"reflect"
)

func GetCanadianVehicleSpecifications(r *CAVehicleSpecRequest) (*BaseApiResponse[GetCAVehicleSpecsResult], error) {
	reqUrl := fmt.Sprintf("%s/GetCanadianVehicleSpecifications?", nhtsaBaseUrl)

	queryParams := url.Values{}

	rV := reflect.ValueOf(*r)
	rT := rV.Type()

	for i := 0; i < rV.NumField(); i++ {
		if rV.Field(i).Interface() != "" {
			queryParams.Add(rT.Field(i).Name, rV.Field(i).Interface().(string))
		}
	}

	queryParams.Add("format", "json")

	res, err := doRequest[BaseApiResponse[GetCAVehicleSpecsResult]](reqUrl + queryParams.Encode())
	if err != nil {
		return nil, fmt.Errorf("error getting CA vehicle specs: %s", err)
	}

	return res, nil
}
