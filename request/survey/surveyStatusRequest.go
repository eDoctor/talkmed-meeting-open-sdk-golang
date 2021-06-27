package survey

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

/**
 * @Description:状态请求
 */
type StatusRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description:
 *
 */
func NewSurveyStatusRequest(surveyId int64) *StatusRequest {

	r := new(StatusRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.SurveyStatusRequestUri, surveyId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 *
 */
func (r *StatusRequest) SetApiVersion(apiVersion string) *StatusRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *StatusRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	paramsByte, _ := json.Marshal(r)

	return r.bRequest.HttpPut(paramsByte)
}
