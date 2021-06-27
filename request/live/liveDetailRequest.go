package live

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type DetailRequest struct {
	bRequest *request.BaseRequest
}

func NewLiveDetailRequest(roomId int64) *DetailRequest {

	r := new(DetailRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveDetailRequestUri, roomId)

	return r
}

func (r *DetailRequest) SetApiVersion(apiVersion string) *DetailRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *DetailRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")
	return resResponse
}
