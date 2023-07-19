package live_speaker

import (
	"encoding/json"
	"fmt"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type UpdateRequest struct {
	bRequest *request.BaseRequest
	Id uint64 `json:"id"`
	SpeakerPassword string `json:"speaker_password"`
}

func NewSpeakerUpdateRequest(roomId int64, id int64) *UpdateRequest {
	r := new(UpdateRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveSpeakerUpdateRequestUri, roomId, id)
	return r
}

func (r *UpdateRequest) SetApiVersion(apiVersion string) *UpdateRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *UpdateRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpPost(registerByte)
}