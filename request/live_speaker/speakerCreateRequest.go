package live_speaker

import (
	"encoding/json"
	"fmt"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type CreateRequest struct {
	bRequest *request.BaseRequest
	Speakers  []map[string]interface{} `json:"speakers"`
}

func NewSpeakerCreateRequest(roomId int64) *CreateRequest {
	r := new(CreateRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveSpeakerCreateRequestUri, roomId)
	return r
}

func (r *CreateRequest) SetApiVersion(apiVersion string) *CreateRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *CreateRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpPost(registerByte)
}