package auth

import (
	"testing"
	"fmt"
)


func TestLogin(t *testing.T) {
	token,err := Login("http://www.unknowname.kl","tkg","123456")
	fmt.Println("err is ",err)
	fmt.Println("token is ",token)
}


func TestGetInput(t *testing.T) {
	u,p := GetInput()
	fmt.Println("username is ",u)
	fmt.Println("password is ",p)
}