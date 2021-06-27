package live

import (
	"encoding/json"
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
)

type AccessRequest struct {
	bRequest     *request.BaseRequest
	UserId       int64  `json:"user_id"`
	RoomRole     uint8  `json:"room_role"`
	RoomPassword string `json:"room_password"`
}

/**
 * @Description:
 * @param roomId
 * @return *AccessRequest
 */
func NewLiveAccessRequest(roomId int64) *AccessRequest {

	r := new(AccessRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveAccessRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *AccessRequest
 */
func (r *AccessRequest) SetApiVersion(apiVersion string) *AccessRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *AccessRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpPost(registerByte)
}
