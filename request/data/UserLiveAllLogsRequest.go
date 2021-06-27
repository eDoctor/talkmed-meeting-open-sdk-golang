package data

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
)

type UserLiveAllLogsRequest struct {
	bRequest *request.BaseRequest
}

func NewUserLiveAllLogsRequest(roomId int64) *UserLiveAllLogsRequest {

	r := new(UserLiveAllLogsRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataUserLiveAllLogsRequestUri, roomId)

	return r
}

func (r *UserLiveAllLogsRequest) SetApiVersion(apiVersion string) *UserLiveAllLogsRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *UserLiveAllLogsRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
