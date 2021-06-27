package live_file

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
)

type FileBindDeleteRequest struct {
	bRequest *request.BaseRequest
	FileIds  []int64 `json:"file_ids"`
}

/**
 * 会议文件上传
 * @Description:
 * @return *FileBindDeleteRequest
 */
func NewLiveFileBindDeleteRequest(FileId int64) *FileBindDeleteRequest {

	r := new(FileBindDeleteRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveFileBindDeleteRequestUri, FileId)

	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FileBindDeleteRequest
 */
func (r *FileBindDeleteRequest) SetApiVersion(apiVersion string) *FileBindDeleteRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FileBindDeleteRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	return r.bRequest.HttpDelete([]byte(""))
}
