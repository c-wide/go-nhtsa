package nhtsa

import (
	"fmt"
	"net/url"
)

func GetAllManufacturers(mType string, page string) (*BaseApiResponse[GetAllManufacturersResult], error) {
	reqUrl := fmt.Sprintf("%s/GetAllManufacturers?", nhtsaBaseUrl)

	queryParams := url.Values{}

	if mType != "" {
		queryParams.Add("ManufacturerType", mType)
	}

	if page != "" {
		queryParams.Add("page", page)
	}

	queryParams.Add("format", "json")

	res, err := doRequest[BaseApiResponse[GetAllManufacturersResult]](reqUrl + queryParams.Encode())
	if err != nil {
		return nil, fmt.Errorf("error getting manufacturers: %s", err)
	}

	return res, nil
}

func GetManufacturerDetails(mfr string) (*BaseApiResponse[GetManufacturerDetailsResult], error) {
	reqUrl := fmt.Sprintf("%s/GetManufacturerDetails/%s?format=json", nhtsaBaseUrl, mfr)

	res, err := doRequest[BaseApiResponse[GetManufacturerDetailsResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting manufacturer details: %s", err)
	}

	return res, nil
}

func GetMakesForManufacturer(mfr string) (*BaseApiResponse[GetMakeManufacturerNameResult], error) {
	reqUrl := fmt.Sprintf("%s/GetMakeForManufacturer/%s?format=json", nhtsaBaseUrl, mfr)

	res, err := doRequest[BaseApiResponse[GetMakeManufacturerNameResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting manufacturer makes: %s", err)
	}

	return res, nil
}

func GetMakesForManufacturerWithYear(mfr string, year string) (*BaseApiResponse[GetMakeManufacturerNameYearResult], error) {
	reqUrl := fmt.Sprintf("%s/GetMakesForManufacturerAndYear/%s?year=%s&format=json", nhtsaBaseUrl, mfr, year)

	res, err := doRequest[BaseApiResponse[GetMakeManufacturerNameYearResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting manufacturer makes: %s", err)
	}

	return res, nil
}
