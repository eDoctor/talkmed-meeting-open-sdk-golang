package survey_question

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

/**
 * @Description:添加调研请求
 */
type CreateRequest struct {
	bRequest     *request.BaseRequest
	Title        string                   `json:"title"`
	QuestionType uint8                    `json:"question_type"`
	IsMust       uint8                    `json:"is_must,omitempty"`
	Introduction string                   `json:"introduction,omitempty"`
	Options      []map[string]interface{} `json:"options,omitempty"`
}

/**
 * @Description:
 *
 */
func NewQuestionCreateRequest(surveyId int64) *CreateRequest {

	r := new(CreateRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.SurveyQuestionCreateRequestUri, surveyId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 *
 */
func (r *CreateRequest) SetApiVersion(apiVersion string) *CreateRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *CreateRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)

	return r.bRequest.HttpPost(registerByte)
}
