package live_fly

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type FlyLayoutRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description:
 * @param roomId
 * @return *FlyLayoutRequest
 */
func NewLiveFlyLayoutRequest(roomId int64) *FlyLayoutRequest {

	r := new(FlyLayoutRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveFlyLayoutRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FlyLayoutRequest
 */
func (r *FlyLayoutRequest) SetApiVersion(apiVersion string) *FlyLayoutRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FlyLayoutRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client
	return r.bRequest.HttpGet("")
}
