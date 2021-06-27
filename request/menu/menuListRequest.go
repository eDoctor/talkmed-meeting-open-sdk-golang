package menu

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type ListRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description:
 * @param roomId
 * @return *MenuListRequest
 */
func NewMenuListRequest(roomId int64) *ListRequest {

	r := new(ListRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.MenuListRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *MenuListRequest
 */
func (r *ListRequest) SetApiVersion(apiVersion string) *ListRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *ListRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
