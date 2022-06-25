package serviceb

import "fmt"

type ServiceB struct {
	Url string
}

type DataAResponse struct {
	A string
	B string
}

func (s ServiceB) GetDataFromB() (DataAResponse, error) {

	fmt.Println("run service b url :", s.Url)

	return DataAResponse{}, nil
}
