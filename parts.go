package nhtsa

import (
	"fmt"
	"net/url"
	"reflect"
)

func GetParts(r *PartsRequest) (*BaseApiResponse[GetPartsResult], error) {
	reqUrl := fmt.Sprintf("%s/GetParts?", nhtsaBaseUrl)

	queryParams := url.Values{}

	rV := reflect.ValueOf(*r)
	rT := rV.Type()

	for i := 0; i < rV.NumField(); i++ {
		if rV.Field(i).Interface() != "" {
			queryParams.Add(rT.Field(i).Name, rV.Field(i).Interface().(string))
		}
	}

	queryParams.Add("format", "json")

	res, err := doRequest[BaseApiResponse[GetPartsResult]](reqUrl + queryParams.Encode())
	if err != nil {
		return nil, fmt.Errorf("error getting parts: %s", err)
	}

	return res, nil
}
