package data

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type UserLiveLogsByRtcTimeRequest struct {
	bRequest *request.BaseRequest
}

func NewUserLiveLogsByRtcTimeRequest(roomId int64) *UserLiveLogsByRtcTimeRequest {

	r := new(UserLiveLogsByRtcTimeRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataUserLiveLogsByRtcTimeRequestUri, roomId)

	return r
}

func (r *UserLiveLogsByRtcTimeRequest) SetApiVersion(apiVersion string) *UserLiveLogsByRtcTimeRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *UserLiveLogsByRtcTimeRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
