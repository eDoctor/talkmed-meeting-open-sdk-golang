package live_fly

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type FlyMemberRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description:
 * @param roomId
 * @return *FlyMemberRequest
 */
func NewLiveFlyMemberRequest(roomId int64) *FlyMemberRequest {

	r := new(FlyMemberRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveFlyMemberRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FlyMemberRequest
 */
func (r *FlyMemberRequest) SetApiVersion(apiVersion string) *FlyMemberRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FlyMemberRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client
	return r.bRequest.HttpGet("")
}
