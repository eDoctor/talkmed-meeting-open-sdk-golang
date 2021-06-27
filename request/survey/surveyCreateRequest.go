package survey

import (
	"encoding/json"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

/**
 * @Description:添加调研请求
 */
type CreateRequest struct {
	bRequest     *request.BaseRequest
	Title        string `json:"title"`
	Introduction string `json:"introduction,omitempty"`
	Banner       int64  `json:"banner,omitempty"`
	Module       int64  `json:"module"`
	ModuleId     int64  `json:"module_id"`
	IsAnswerShow uint8  `json:"is_answer_show,omitempty"`
}

/**
 * @Description:
 *
 */
func NewSurveyCreateRequest() *CreateRequest {

	r := new(CreateRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.SurveyCreateRequestUri
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
