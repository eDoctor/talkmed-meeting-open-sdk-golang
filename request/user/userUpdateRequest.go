package user

import (
	"encoding/json"
	"fmt"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type UpdateRequest struct {
	bRequest *request.BaseRequest
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar,omitempty"`
	UserRole uint8 `json:"user_role,omitempty"`
	RealName string `json:"realname,omitempty"`
	Email    string `json:"email,omitempty"`
}

func NewUserUpdateRequest(userId int64) *UpdateRequest {
	r := new(UpdateRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.UserUpdateRequestUri, userId)
	return r
}

func (r *UpdateRequest) SetApiVersion(apiVersion string) *UpdateRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *UpdateRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)

	fmt.Printf("%v\n", registerByte)
	return r.bRequest.HttpPut(registerByte)
}
