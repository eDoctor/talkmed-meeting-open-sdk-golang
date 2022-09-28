package user

import (
	"encoding/json"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

const RegisterLoginUri = "/open/register"

/**
 * @Description:
 */
type LoginRequest struct {
	bRequest     *request.BaseRequest
	Type         string `json:"type,omitempty"`
	Login        uint8  `json:"login,omitempty"`
	Platform     string `json:"platform"`
	Nickname     string `json:"nickname"`
	Email        string `json:"email,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	Realname     string `json:"realname,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Password     string `json:"password"`
	UserRole     uint8  `json:"user_role,omitempty"`
	Openid       string `json:"openid,omitempty"`
	Unionid      string `json:"unionid,omitempty"`
	RoomId       string `json:"room_id,omitempty"`
	RoomRole     uint8  `json:"room_role,omitempty"`
	RoomPassword string `json:"room_password,omitempty"`
}

/**
 * @Description:
 *
 */
func NewLoginRequest() *LoginRequest {

	r := new(LoginRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.UserLoginRequestUri

	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 *
 */
func (r *LoginRequest) SetApiVersion(apiVersion string) *LoginRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description:
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *LoginRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)

	return r.bRequest.HttpPost(registerByte)
}
