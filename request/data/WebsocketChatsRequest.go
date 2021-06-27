package data

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type WebsocketChatsRequest struct {
	bRequest *request.BaseRequest
}

func NewWebsocketChatsRequest(roomId int64) *WebsocketChatsRequest {

	r := new(WebsocketChatsRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataWebsocketChatsRequestUri, roomId)

	return r
}

func (r *WebsocketChatsRequest) SetApiVersion(apiVersion string) *WebsocketChatsRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *WebsocketChatsRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpGet("")

	return resResponse
}
