package survey

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
	Title        string `json:"title,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Banner       int64  `json:"banner,omitempty"`
	Module       int64  `json:"module"`
	ModuleId     int64  `json:"module_id"`
	IsShow       uint8  `json:"is_show,omitempty"`
	IsAnswerShow uint8  `json:"is_answer_show,omitempty"`
}

/**
 * @Description:
 *
 */
func NewSurveyUpdateRequest(surveyId int64) *UpdateRequest {

	r := new(UpdateRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.SurveyUpdateRequestUri, surveyId)
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
	return r.bRequest.HttpPut(registerByte)
}
