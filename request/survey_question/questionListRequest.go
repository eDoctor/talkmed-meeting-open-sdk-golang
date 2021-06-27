package survey_question

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type QuestionListRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description:
 * @param roomId
 * @return *QuestionListRequest
 */
func NewQuestionListRequest(surveyId int64) *QuestionListRequest {

	r := new(QuestionListRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.SurveyQuestionListRequestUri, surveyId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *QuestionListRequest
 */
func (r *QuestionListRequest) SetApiVersion(apiVersion string) *QuestionListRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *QuestionListRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")
	return resResponse
}
