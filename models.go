package nhtsa

import "fmt"

func GetModelsForMake(mName string) (*BaseApiResponse[GetModelsResult], error) {
	reqUrl := fmt.Sprintf("%s/GetModelsForMake/%s?format=json", nhtsaBaseUrl, mName)

	res, err := doRequest[BaseApiResponse[GetModelsResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting models: %s", err)
	}

	return res, nil
}
