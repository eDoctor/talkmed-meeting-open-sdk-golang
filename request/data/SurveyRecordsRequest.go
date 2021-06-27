package data

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type SurveyRecordsRequest struct {
	bRequest *request.BaseRequest
}

func NewSurveyRecordsRequest(roomId int64) *SurveyRecordsRequest {

	r := new(SurveyRecordsRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataSurveyRecordsRequestUri, roomId)
	return r
}

func (r *SurveyRecordsRequest) SetApiVersion(apiVersion string) *SurveyRecordsRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *SurveyRecordsRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
