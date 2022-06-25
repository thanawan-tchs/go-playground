package servicea

import (
	"encoding/json"
	"go-gin-playground/pkg/https"
)

type ServiceA struct {
	Url string
}

type DataARequest struct {
	A string `json:"value_a"`
	B string `json:"value_b"`
}

type DataAResponse struct {
	C string `json:"value_c"`
	D string `json:"value_d"`
}

func header() map[string]string {
	return map[string]string{
		"content-type": "application/json",
		"accept":       "application/json",
	}
}

func (s ServiceA) GetDataFromA(body DataARequest) (DataAResponse, error) {

	rqByte, err := json.Marshal(&body)
	if err != nil {
		return DataAResponse{}, err
	}

	rq := https.HttpRequest{
		Header: header(),
		Method: "GET",
		Url:    s.Url + "/service/a",
		Body:   rqByte,
	}

	resByte, err := https.SendHttpRequest(rq)
	if err != nil {
		return DataAResponse{}, err
	}

	var resp DataAResponse
	err = json.Unmarshal(resByte, &resp)
	if err != nil {
		return DataAResponse{}, err
	}

	return resp, nil
}
