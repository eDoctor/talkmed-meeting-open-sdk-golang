package live_white

import (
	"encoding/json"
	"fmt"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type DeleteRequest struct {
	bRequest *request.BaseRequest
	Type    int8 `json:"type"`
	All    int8 `json:"all"`
	Content string `json:"content,omitempty"`
}

func NewWhiteDeleteRequest(roomId int64) *DeleteRequest {
	r := new(DeleteRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveWhiteDeleteRequestUri, roomId)
	return r
}

func (r *DeleteRequest) SetApiVersion(apiVersion string) *DeleteRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *DeleteRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpDelete(registerByte)
}