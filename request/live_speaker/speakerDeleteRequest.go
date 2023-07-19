package live_speaker

import (
	"encoding/json"
	"fmt"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type DeleteRequest struct {
	bRequest *request.BaseRequest
}

func NewSpeakerDeleteRequest(roomId int64, id int64) *DeleteRequest {
	r := new(DeleteRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveSpeakerDeleteRequestUri, roomId, id)
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