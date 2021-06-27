package data

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
)

type AccessRecordsRequestRequest struct {
	bRequest *request.BaseRequest
}

func NewAccessRecordsRequestRequest(roomId int64) *AccessRecordsRequestRequest {

	r := new(AccessRecordsRequestRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataAccessRecordsRequestUri, roomId)

	return r
}

func (r *AccessRecordsRequestRequest) SetApiVersion(apiVersion string) *AccessRecordsRequestRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *AccessRecordsRequestRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
