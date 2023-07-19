package live_speaker

import (
	"fmt"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type DetailRequest struct {
	bRequest *request.BaseRequest
}

func NewSpeakerDetailRequest(roomId int64, id int64) *DetailRequest {
	r := new(DetailRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveSpeakerDetailRequestUri, roomId, id)
	return r
}

func (r *DetailRequest) SetApiVersion(apiVersion string) *DetailRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *DetailRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client
	//registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpGet("")
}