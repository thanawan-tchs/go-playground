package https

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpRequest struct {
	Method string
	Url    string
	Body   []byte
	Params string
	Header map[string]string
}

func SendHttpRequest(rq HttpRequest) ([]byte, error) {

	req, err := http.NewRequest(rq.Method, rq.Url+rq.Params, bytes.NewReader(rq.Body))
	if err != nil {
		return []byte{}, err
	}

	for k, v := range rq.Header {
		req.Header.Set(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return resBody, nil
}
