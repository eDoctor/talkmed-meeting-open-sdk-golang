package user

import (
	"fmt"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
	"net/url"
)

type AuthRequest struct {
	bRequest *request.BaseRequest
	AuthToken string `json:"auth_token"`
}

func NewUserAuthRequest() *AuthRequest {
	r := new(AuthRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.UserAuthRequestUri
	return r
}

func (r *AuthRequest) SetApiVersion(apiVersion string) *AuthRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *AuthRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client

	queryStringStruct := url.Values{}
	queryStringStruct.Set("auth_token", fmt.Sprintf("%s", r.AuthToken))
	queryStr := queryStringStruct.Encode()

	resResponse := r.bRequest.HttpGet(queryStr)

	return resResponse
}


