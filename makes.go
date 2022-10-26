package nhtsa

import (
	"fmt"
)

func GetAllMakes() (*BaseApiResponse[GetAllMakesResult], error) {
	reqUrl := fmt.Sprintf("%s/getallmakes?format=json", nhtsaBaseUrl)

	res, err := doRequest[BaseApiResponse[GetAllMakesResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting all makes: %s", err)
	}

	return res, nil
}
