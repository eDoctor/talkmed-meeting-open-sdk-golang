package live

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

/**
 * @Description:
 */
type DeleteRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description: 会议删除
 * @param roomId
 * @return *DeleteRequest
 */
func NewLiveDeleteRequest(roomId int64) *DeleteRequest {

	r := new(DeleteRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveDetailRequestUri, roomId)

	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *DeleteRequest
 */
func (r *DeleteRequest) SetApiVersion(apiVersion string) *DeleteRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *DeleteRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	resResponse := r.bRequest.HttpDelete([]byte(""))
	return resResponse
}
