package live_file

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type FileBindRequest struct {
	bRequest *request.BaseRequest
	FileIds  []int64 `json:"file_ids"`
}

/**
 * 会议文件绑定
 * @Description:
 * @return *FileBindRequest
 */
func NewLiveFileBindRequest(roomId int64) *FileBindRequest {

	r := new(FileBindRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveFileBindRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FileBindRequest
 */
func (r *FileBindRequest) SetApiVersion(apiVersion string) *FileBindRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FileBindRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	body, _ := json.Marshal(r)

	return r.bRequest.HttpPost(body)
}
