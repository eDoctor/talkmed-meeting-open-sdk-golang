package survey_question

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
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
func NewQuestionDetailRequest(surveyId int64, questionId int64) *StatusRequest {

	r := new(StatusRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.SurveyQuestionDetailRequestUri, surveyId, questionId)
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
	resResponse := r.bRequest.HttpGet("")
	return resResponse
}
