package gateway

import (
	"fmt"
	servicea "go-gin-playground/gateway/serviceA"
	serviceb "go-gin-playground/gateway/serviceB"
)

var gw *externalHttp

type externalHttp struct {
	A servicea.ServiceA
	B serviceb.ServiceB
}

func ServiceA() servicea.ServiceA {
	return gw.A
}

func ServiceB() serviceb.ServiceB {
	return gw.B
}

func InitGateway() {

	gw = &externalHttp{
		A: servicea.ServiceA{
			Url: "http://localhost:8081",
		},
		B: serviceb.ServiceB{
			Url: "http://localhost:8082",
		},
	}

	fmt.Println("---- init gateway ----")
	fmt.Println("---- a :", gw.A.Url)
	fmt.Println("---- b :", gw.B.Url)
	fmt.Println("----------------------")
}
