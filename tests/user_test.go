package tests

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk/request/user"
	"testing"
)

/**
 * @Description: 注册
 * @param t
 */
func TestRegister(t *testing.T) {

	client := getClient()
	userRequest := user.NewRegisterRequest()
	userRequest.Email = "11974097@qq.com"
	userRequest.Nickname = "yzh"
	userRequest.Platform = "open"
	userRequest.Password = "123456"
	userRequest.Login = 1

	res := userRequest.Request(client)
	fmt.Println(res)
}

/**
 * @Description: 登录
 * @param t
 */
func TestLogin(t *testing.T) {

	client := getClient()
	userRequest := user.NewLoginRequest()
	userRequest.Email = "11974097@qq.com"
	userRequest.Nickname = "yzh"
	userRequest.Platform = "open"
	userRequest.Password = "123456"
	userRequest.Login = 1

	res := userRequest.Request(client)
	fmt.Println(res)
}
