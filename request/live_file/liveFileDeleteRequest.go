package live_file

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
)

type FileDeleteRequest struct {
	bRequest *request.BaseRequest
}

/**
 * 会议文件上传
 * @Description:
 * @return *FileDeleteRequest
 */
func NewLiveFileDeleteRequest(FileId int64) *FileDeleteRequest {

	r := new(FileDeleteRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveFileDeleteRequestUri, FileId)

	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FileDeleteRequest
 */
func (r *FileDeleteRequest) SetApiVersion(apiVersion string) *FileDeleteRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FileDeleteRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	return r.bRequest.HttpDelete([]byte(""))
}
