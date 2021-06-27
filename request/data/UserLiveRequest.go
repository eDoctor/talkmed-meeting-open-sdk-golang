package data

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type UserLiveRequest struct {
	bRequest *request.BaseRequest
}

func NewUserLiveRequest(roomId int64) *UserLiveRequest {

	r := new(UserLiveRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataUserLiveRequestUri, roomId)

	return r
}

func (r *UserLiveRequest) SetApiVersion(apiVersion string) *UserLiveRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *UserLiveRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
