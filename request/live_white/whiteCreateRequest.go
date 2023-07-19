package live_white

import (
	"encoding/json"
	"fmt"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type CreateRequest struct {
	bRequest *request.BaseRequest
	Type    int8 `json:"type"`
	Content string `json:"content"`
}

func NewWhiteCreateRequest(roomId int64) *CreateRequest {
	r := new(CreateRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveWhiteCreateRequestUri, roomId)
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