package live

import (
	"encoding/json"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
)

/**
 * @Description:用户注册请求
 */
type CreateRequest struct {
	bRequest             *request.BaseRequest
	Title                string                   `json:"title"`
	StartAt              string                   `json:"start_at"`
	EndAt                string                   `json:"end_at"`
	LiveType             uint8                    `json:"live_type,omitempty"`
	Introduction         string                   `json:"introduction,omitempty"`
	Banner               string                   `json:"banner,omitempty"`
	ExtendInfo           map[string]interface{}   `json:"extend_info,omitempty"`
	UserId               int64                    `json:"user_id,omitempty"`
	IsSpeakerOnePassword uint8                    `json:"is_speaker_one_password"`
	Speakers             []map[string]interface{} `json:"speakers,omitempty"`
	AccessType           int64                    `json:"access_type"`
	WatchPassword        string                   `json:"watch_password,omitempty"`
	StreamProvider       uint8                    `json:"stream_provider,omitempty"`
}

/**
 * @Description:
 *
 */
func NewCreateRequest() *CreateRequest {

	r := new(CreateRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.LiveCreateRequestUri

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
