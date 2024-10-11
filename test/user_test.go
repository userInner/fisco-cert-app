package test

import (
	"fisco-cert-app/logic"
	"fisco-cert-app/models"
	"fmt"
	"testing"
)

func TestSignUp(t *testing.T) {
	user := &models.ParamSignUp{
		Username:   "ff",
		Password:   "123456",
		RePassword: "123456",
	}

	err := logic.SignUp(user)
	if err != nil {
		fmt.Printf("Test Logic SignUp failed,err: %v", err.Error())
		return
	}
	fmt.Println("Test Logic SignUp Success")
}
