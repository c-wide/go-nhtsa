package nhtsa

import (
	"fmt"
	"net/url"
	"reflect"
)

func GetParts(r *PartsRequest) (*BaseApiResponse[GetPartsResult], error) {
	reqUrl := fmt.Sprintf("%s/GetParts?", nhtsaBaseUrl)

	queryParams := url.Values{}

	v := reflect.ValueOf(r)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() != "" {
			queryParams.Add(typeOfS.Field(i).Name, v.Field(i).Interface().(string))
		}
	}

	queryParams.Add("format", "json")

	res, err := doRequest[BaseApiResponse[GetPartsResult]](reqUrl + queryParams.Encode())
	if err != nil {
		return nil, fmt.Errorf("error getting parts: %s", err)
	}

	return res, nil
}
