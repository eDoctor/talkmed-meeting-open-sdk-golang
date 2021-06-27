package live_file

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

type FileBindListRequest struct {
	bRequest *request.BaseRequest
	FileIds  []int64 `json:"file_ids"`
}

/**
 * 会议文件绑定
 * @Description:
 * @return *FileBindListRequest
 */
func NewLiveFileBindListRequest(roomId int64) *FileBindListRequest {

	r := new(FileBindListRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveFileBindListRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FileBindListRequest
 */
func (r *FileBindListRequest) SetApiVersion(apiVersion string) *FileBindListRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FileBindListRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client

	return r.bRequest.HttpGet("")
}
