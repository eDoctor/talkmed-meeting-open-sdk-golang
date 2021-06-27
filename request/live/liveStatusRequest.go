package live

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type StatusRequest struct {
	bRequest *request.BaseRequest
	Status   uint8 `json:"status"`
}

/**
 * @Description:
 * @param roomId
 * @return *StatusRequest
 */
func NewLiveStatusRequest(roomId int64) *StatusRequest {

	r := new(StatusRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveStatusRequestUri, roomId)
	r.Status = 1 //默认值
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *StatusRequest
 */
func (r *StatusRequest) SetApiVersion(apiVersion string) *StatusRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *StatusRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpPost(registerByte)
}
