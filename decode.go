package nhtsa

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func DecodeVin(request *VinRequest) (*DecodeResponse[DecodeResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevin/%s?format=json&modelyear=%s", nhtsaBaseUrl, request.Vin, request.Year)

	res, err := doRequest[DecodeResponse[DecodeResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinFlat(request *VinRequest) (*DecodeResponse[DecodeFlatResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvalues/%s?format=json&modelyear=%s", nhtsaBaseUrl, request.Vin, request.Year)

	res, err := doRequest[DecodeResponse[DecodeFlatResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinExtended(request *VinRequest) (*DecodeResponse[DecodeResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvaluesextended/%s?format=json&modelyear=%s", nhtsaBaseUrl, request.Vin, request.Year)

	res, err := doRequest[DecodeResponse[DecodeResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinFlatExtended(request *VinRequest) (*DecodeResponse[DecodeFlatResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvaluesextended/%s?format=json&modelyear=%s", nhtsaBaseUrl, request.Vin, request.Year)

	res, err := doRequest[DecodeResponse[DecodeFlatResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeWmi(wmi string) (*DecodeResponse[DecodeWmiResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodewmi/%s?format=json", nhtsaBaseUrl, wmi)

	res, err := doRequest[DecodeResponse[DecodeWmiResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding WMI: %s", err)
	}

	return res, nil
}

// https://vpic.nhtsa.dot.gov/api/home/index/faq
//
// per above FAQ, it's best to use standard decode
// APIs when processing large amounts of VINs.
func DecodeVinBatch(vList *[]VinRequest) (*[]DecodeFlatResponseResult, error) {
	g := new(errgroup.Group)

	retList := make([]DecodeFlatResponseResult, len(*vList))

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
