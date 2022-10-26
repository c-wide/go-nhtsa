package nhtsa

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func DecodeVin(request *VinRequest) (*BaseApiResponse[DecodeVinResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevin/%s?format=json&modelyear=%s", nhtsaBaseUrl, request.Vin, request.Year)

	res, err := doRequest[BaseApiResponse[DecodeVinResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinFlat(request *VinRequest) (*BaseApiResponse[DecodeFlatResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvalues/%s?format=json&modelyear=%s", nhtsaBaseUrl, request.Vin, request.Year)

	res, err := doRequest[BaseApiResponse[DecodeFlatResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinExtended(request *VinRequest) (*BaseApiResponse[DecodeVinResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvaluesextended/%s?format=json&modelyear=%s", nhtsaBaseUrl, request.Vin, request.Year)

	res, err := doRequest[BaseApiResponse[DecodeVinResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinFlatExtended(request *VinRequest) (*BaseApiResponse[DecodeFlatResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvaluesextended/%s?format=json&modelyear=%s", nhtsaBaseUrl, request.Vin, request.Year)

	res, err := doRequest[BaseApiResponse[DecodeFlatResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeWmi(wmi string) (*BaseApiResponse[DecodeWmiResult], error) {
	reqUrl := fmt.Sprintf("%s/decodewmi/%s?format=json", nhtsaBaseUrl, wmi)

	res, err := doRequest[BaseApiResponse[DecodeWmiResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding WMI: %s", err)
	}

	return res, nil
}

// https://vpic.nhtsa.dot.gov/api/home/index/faq
//
// per above FAQ, it's best to use standard decode
// APIs when processing large amounts of VINs.
func DecodeVinBatch(vList *[]VinRequest) (*[]DecodeFlatResult, error) {
	g := new(errgroup.Group)

	retList := make([]DecodeFlatResult, len(*vList))

	for vIdx, vInfo := range *vList {
		vIdx := vIdx
		vInfo := vInfo

		g.Go(func() error {
			res, err := DecodeVinFlat(&vInfo)
			if err != nil {
				return err
			}

			retList[vIdx] = res.Results[0]

			return nil
		})

		time.Sleep(500 * time.Millisecond)
	}

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("error performing requests. %s", err)
	}

	return &retList, nil
}
