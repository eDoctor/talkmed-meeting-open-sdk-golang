package tests

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request/user"
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

func TestUserUpdate(t *testing.T) {
	c := getClient()
	u := user.NewUserUpdateRequest(46847)
	u.Nickname = "simplevvv"
	r := u.Request(c)
	fmt.Println(r)
}

func TestUserDelete(t *testing.T) {
	c := getClient()
	u := user.NewUserDeleteRequest(45030)
	r := u.Request(c)
	fmt.Println(r)
}

func TestUserAuth(t *testing.T) {
	c := getClient()
	u := user.NewUserAuthRequest()
	u.AuthToken = "363f9c94-3e1c-5740-f85c-cfc4faeea7fd"
	r := u.Request(c)
	fmt.Println(r)
}
