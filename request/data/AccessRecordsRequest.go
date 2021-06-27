package data

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type AccessRecordsRequestRequest struct {
	bRequest *request.BaseRequest
}

func NewAccessRecordsRequestRequest(roomId int64) *AccessRecordsRequestRequest {

	r := new(AccessRecordsRequestRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataAccessRecordsRequestUri, roomId)

	return r
}

func (r *AccessRecordsRequestRequest) SetApiVersion(apiVersion string) *AccessRecordsRequestRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *AccessRecordsRequestRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
