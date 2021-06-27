package survey_question

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

/**
 * @Description:添加调研请求
 */
type DeleteRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description:
 *
 */
func NewQuestionDeleteRequest(surveyId int64, questionId int64) *DeleteRequest {

	r := new(DeleteRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.SurveyQuestionDeleteRequestUri, surveyId, questionId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 *
 */
func (r *DeleteRequest) SetApiVersion(apiVersion string) *DeleteRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *DeleteRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	return r.bRequest.HttpDelete([]byte(""))
}
