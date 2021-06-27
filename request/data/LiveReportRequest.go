package data

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type LiveReportRequest struct {
	bRequest *request.BaseRequest
}

func NewLiveReportRequest(roomId int64) *LiveReportRequest {

	r := new(LiveReportRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataLiveReportRequestUri, roomId)

	return r
}

func (r *LiveReportRequest) SetApiVersion(apiVersion string) *LiveReportRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *LiveReportRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
