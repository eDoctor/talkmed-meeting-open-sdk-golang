package live

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

/**
 * @Description:会议修改请求
 */
type UpdateRequest struct {
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
func NewUpdateRequest(roomId int64) *UpdateRequest {

	r := new(UpdateRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveUpdateRequestUri, roomId)
	r.initStruct()

	return r
}

/**
 * @Description: todo... 初始化 int 类型 带有零值
 * @receiver r
 */
func (r *UpdateRequest) initStruct() {
	r.IsSpeakerOnePassword = 1
	r.AccessType = 0
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
