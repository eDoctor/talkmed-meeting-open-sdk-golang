package live_chat

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type ChatCheckRequest struct {
	bRequest    *request.BaseRequest
	AuditStatus uint8 `json:"audit_status"`
}

/**
 * 聊天审核
 * @Description:
 * @param roomId
 * @param chatId
 * @return *ChatCheckRequest
 */
func NewLiveChatCheckRequest(roomId int64, chatId int64) *ChatCheckRequest {

	r := new(ChatCheckRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveChatCheckRequestUri, roomId, chatId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *ChatCheckRequest
 */
func (r *ChatCheckRequest) SetApiVersion(apiVersion string) *ChatCheckRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *ChatCheckRequest) Request(client *meeting.MeetClient) request.ApiResponse {
	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpPut(registerByte)
}
