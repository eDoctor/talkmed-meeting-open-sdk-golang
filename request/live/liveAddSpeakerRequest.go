package live

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type AddSpeakerRequest struct {
	bRequest *request.BaseRequest
	Speakers []map[string]interface{} `json:"speakers"`
}

/**
 * @Description:
 * @param roomId
 * @return *AddSpeakerRequest
 */
func NewLiveAddSpeakerRequest(roomId int64) *AddSpeakerRequest {

	r := new(AddSpeakerRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveAddSpeakerRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *AddSpeakerRequest
 */
func (r *AddSpeakerRequest) SetApiVersion(apiVersion string) *AddSpeakerRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *AddSpeakerRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpPost(registerByte)
}
