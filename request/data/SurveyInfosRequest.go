package data

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
)

type SurveyInfosRequest struct {
	bRequest *request.BaseRequest
}

func NewSurveyInfosRequest(roomId int64) *SurveyInfosRequest {

	r := new(SurveyInfosRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataSurveyInfosRequestUri, roomId)

	return r
}

func (r *SurveyInfosRequest) SetApiVersion(apiVersion string) *SurveyInfosRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *SurveyInfosRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
