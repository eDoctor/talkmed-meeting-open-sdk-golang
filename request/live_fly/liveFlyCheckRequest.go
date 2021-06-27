package live_fly

import (
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type FlyCheckRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description:
 * @param roomId
 * @return *FlyCheckRequest
 */
func NewLiveFlyCheckRequest() *FlyCheckRequest {

	r := new(FlyCheckRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.LiveFlyCheckRequestUri
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FlyCheckRequest
 */
func (r *FlyCheckRequest) SetApiVersion(apiVersion string) *FlyCheckRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FlyCheckRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client
	return r.bRequest.HttpGet("")
}
