package data

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type AboutLiveRequest struct {
	bRequest *request.BaseRequest
}

func NewAboutLiveRequest(surveyId int64) *AboutLiveRequest {

	r := new(AboutLiveRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataAboutLiveRequestUri, surveyId)
	return r
}

func (r *AboutLiveRequest) SetApiVersion(apiVersion string) *AboutLiveRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *AboutLiveRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
