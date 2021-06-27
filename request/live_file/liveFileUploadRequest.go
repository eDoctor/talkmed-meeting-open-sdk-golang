package live_file

import (
	"encoding/json"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

type FileUploadRequest struct {
	bRequest *request.BaseRequest
	FilePath string
	RoomId   int64 `json:"room_id"`
	UserId   int64 `json:"user_id"`
}

/**
 * 会议文件上传
 * @Description:
 * @return *FileUploadRequest
 */
func NewLiveFileUploadRequest() *FileUploadRequest {

	r := new(FileUploadRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.LiveFileUploadRequestUri

	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FileUploadRequest
 */
func (r *FileUploadRequest) SetApiVersion(apiVersion string) *FileUploadRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FileUploadRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client

	registerByte, _ := json.Marshal(r)
	File := map[string]string{"file": r.FilePath}

	return r.bRequest.HttpUpload(File, registerByte)
}
