package nhtsa

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func DecodeVin(vin string, year string) (*NHTSADecodeResponse[NHTSADecodeResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevin/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinFlat(vin string, year string) (*NHTSADecodeResponse[NHTSADecodeFlatResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvalues/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeFlatResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinExtended(vin string, year string) (*NHTSADecodeResponse[NHTSADecodeResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvaluesextended/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeVinFlatExtended(vin string, year string) (*NHTSADecodeResponse[NHTSADecodeFlatResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodevinvaluesextended/%s?format=json&modelyear=%s", nhtsaBaseUrl, vin, year)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeFlatResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding VIN: %s", err)
	}

	return res, nil
}

func DecodeWmi(wmi string) (*NHTSADecodeResponse[NHTSADecodeWmiResponseResult], error) {
	reqUrl := fmt.Sprintf("%s/decodewmi/%s?format=json", nhtsaBaseUrl, wmi)

	res, err := doRequest[NHTSADecodeResponse[NHTSADecodeWmiResponseResult]](reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error decoding WMI: %s", err)
	}

	return res, nil
}

// https://vpic.nhtsa.dot.gov/api/home/index/faq
//
// per above FAQ, it's best to use standard decode APIs
// when processing large amounts of VINs.
//
// TODO: array of tuples or array of VinInfo?
// TODO: Normal, Flat, Extended, ExtendedFlat?
// TODO: Rate limiting?
//
// func DecodeVinBatch(vinList *[][]string) (*[]NHTSADecodeResponse[NHTSADecodeResponseResult], error) {
// 	g := new(errgroup.Group)
//
// 	for vinIdx, vinInfo := range *data {
// 		vinIdx := vinIdx
// 		vinInfo := vinInfo
//
// 		g.Go(func() error {
// 			(*data)[vinIdx].Result = &apiResponse.Results[0]
//
// 			return nil
// 		})
//
// 		time.Sleep(100 * time.Millisecond)
// 	}
//
// 	if err := g.Wait(); err != nil {
// 		return fmt.Errorf("error performing requests. %s", err)
// 	}
//
// 	return nil
// }
