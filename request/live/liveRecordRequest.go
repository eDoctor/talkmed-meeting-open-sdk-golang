package live

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
	"net/url"
)

type RecordRequest struct {
	bRequest *request.BaseRequest
	All      uint8 `json:"all"`
}

/**
 * @Description:
 * @param roomId
 * @return *RecordRequest
 */
func NewLiveRecordRequest(roomId int64) *RecordRequest {

	r := new(RecordRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveRecordRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *RecordRequest
 */
func (r *RecordRequest) SetApiVersion(apiVersion string) *RecordRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *RecordRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	Url := url.Values{}
	Url.Set("all", fmt.Sprintf("%d", r.All))
	return r.bRequest.HttpGet(Url.Encode())
}
