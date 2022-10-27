package nhtsa

import "fmt"

func GetModelsForMakeByName(mName string) (*BaseApiResponse[GetModelsResult], error) {
	reqUrl := fmt.Sprintf("%s/GetModelsForMake/%s?format=json", nhtsaBaseUrl, mName)

	res, err := doRequest[BaseApiResponse[GetModelsResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting models: %s", err)
	}

	return res, nil
}

func GetModelsForMakeById(mId string) (*BaseApiResponse[GetModelsResult], error) {
	reqUrl := fmt.Sprintf("%s/GetModelsForMakeId/%s?format=json", nhtsaBaseUrl, mId)

	res, err := doRequest[BaseApiResponse[GetModelsResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting models: %s", err)
	}

	return res, nil
}
