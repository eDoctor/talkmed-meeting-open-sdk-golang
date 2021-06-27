package data

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type UserLiveUniqueLogsByRtcTimeRequest struct {
	bRequest *request.BaseRequest
}

func NewUserLiveUniqueLogsByRtcTimeRequest(roomId int64) *UserLiveUniqueLogsByRtcTimeRequest {

	r := new(UserLiveUniqueLogsByRtcTimeRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataUserLiveUniqueLogsByRtcTimeRequestUri, roomId)

	return r
}

func (r *UserLiveUniqueLogsByRtcTimeRequest) SetApiVersion(apiVersion string) *UserLiveUniqueLogsByRtcTimeRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *UserLiveUniqueLogsByRtcTimeRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")
	return resResponse
}
