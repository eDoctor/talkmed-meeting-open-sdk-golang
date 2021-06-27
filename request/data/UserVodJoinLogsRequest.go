package data

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
)

type UserVodJoinLogsRequest struct {
	bRequest *request.BaseRequest
}

func NewUserVodJoinLogsRequest(roomId int64) *UserVodJoinLogsRequest {

	r := new(UserVodJoinLogsRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataUserVodJoinLogsRequestUri, roomId)

	return r
}

func (r *UserVodJoinLogsRequest) SetApiVersion(apiVersion string) *UserVodJoinLogsRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *UserVodJoinLogsRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
