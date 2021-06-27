package live_chat

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
	"net/url"
)

type ChatListRequest struct {
	bRequest *request.BaseRequest
	Page     uint8 `json:"page"`
	PageSize uint8 `json:"page_size"`
}

/**
 * @Description:
 * @param roomId
 * @return *ChatListRequest
 */
func NewLiveChatListRequest(roomId int64) *ChatListRequest {

	r := new(ChatListRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveChatListRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *ChatListRequest
 */
func (r *ChatListRequest) SetApiVersion(apiVersion string) *ChatListRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *ChatListRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client

	UrlQuery := url.Values{}
	UrlQuery.Set("page", fmt.Sprintf("%d", r.Page))
	UrlQuery.Set("page_size", fmt.Sprintf("%d", r.PageSize))

	return r.bRequest.HttpGet(UrlQuery.Encode())
}
