package live_speaker

import (
	"fmt"
	meeting "github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
	"net/url"
)

type ListRequest struct {
	bRequest *request.BaseRequest
	Page     uint64
	PageSize uint64
}

func NewSpeakerListRequest(roomId int64) *ListRequest {
	r := new(ListRequest)
	br := new(request.BaseRequest)
	r.bRequest = br
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveSpeakerListRequestUri, roomId)
	return r
}

func (r *ListRequest) SetApiVersion(apiVersion string) *ListRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *ListRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client

	queryStringStruct := url.Values{}
	if r.Page > 0 {
		queryStringStruct.Set("page", fmt.Sprintf("%d", r.Page))
	}

	if r.PageSize > 0 {
		queryStringStruct.Set("page_size", fmt.Sprintf("%d", r.PageSize))
	}

	queryStr := queryStringStruct.Encode()
	resResponse := r.bRequest.HttpGet(queryStr)

	return resResponse
}