package user

import (
	"fmt"
	"go-gin-playground/gateway"
	servicea "go-gin-playground/gateway/serviceA"
)

type User struct {
	Id    string
	Name  string
	Email string
}

func GetUser() (User, error) {
	// connect db, gateway
	u := User{
		Id:    "1234",
		Name:  "hello",
		Email: "aaa@xxx.com",
	}
	// call external A
	dataA, err := gateway.ServiceA().GetDataFromA(createARequest())
	if err != nil {
		return User{}, err
	}

	// call external B
	dataB, err := gateway.ServiceB().GetDataFromB()
	if err != nil {
		return User{}, err
	}

	// connect DB

	fmt.Println("==== data A :", dataA)
	fmt.Println("==== data B :", dataB)
	return u, nil
}

func createARequest() servicea.DataARequest {
	return servicea.DataARequest{
		A: "a",
		B: "b",
	}
}
