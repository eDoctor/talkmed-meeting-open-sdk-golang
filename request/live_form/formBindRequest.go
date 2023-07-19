package live_form

import (
	"encoding/json"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type BindFormRequest struct {
	bRequest *request.BaseRequest
	FormId uint64 `json:"form_id"`
	RoomId uint64 `json:"room_id"`
	CollectionRule  uint8 `json:"collection_rule,omitempty"`
}

func NewLiveBindFormRequest() *BindFormRequest {
	r := new(BindFormRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.LiveBindFormRequestUri
	return r
}

func (r *BindFormRequest) SetApiVersion(apiVersion string) *BindFormRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *BindFormRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpPost(registerByte)
}