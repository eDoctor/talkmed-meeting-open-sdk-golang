package data

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type LiveBaseRequest struct {
	bRequest *request.BaseRequest
}

func NewLiveBaseRequest(roomId int64) *LiveBaseRequest {

	r := new(LiveBaseRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataLiveBaseRequestUri, roomId)

	return r
}

func (r *LiveBaseRequest) SetApiVersion(apiVersion string) *LiveBaseRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *LiveBaseRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client

	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
