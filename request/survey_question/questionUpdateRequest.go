package survey_question

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

/**
 * @Description:添加调研请求
 */
type UpdateRequest struct {
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
func NewQuestionUpdateRequest(surveyId int64, questionId int64) *UpdateRequest {

	r := new(UpdateRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.SurveyQuestionUpdateRequestUri, surveyId, questionId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 *
 */
func (r *UpdateRequest) SetApiVersion(apiVersion string) *UpdateRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *UpdateRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	fmt.Print(string(registerByte))
	return r.bRequest.HttpPut(registerByte)
}
