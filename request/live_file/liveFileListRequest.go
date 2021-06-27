package live_file

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
	"net/url"
)

type FileListRequest struct {
	bRequest *request.BaseRequest
	Page     uint8
	PageSize uint8
	Type     uint8
}

/**
 * @Description:
 * @param roomId
 * @return *FileListRequest
 */
func NewLiveFileListRequest() *FileListRequest {

	r := new(FileListRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.LiveFileListRequestUri
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FileListRequest
 */
func (r *FileListRequest) SetApiVersion(apiVersion string) *FileListRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FileListRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client

	queryStringStruct := url.Values{}
	if r.Page > 0 {
		queryStringStruct.Set("page", fmt.Sprintf("%d", r.Page))
	}

	if r.PageSize > 0 {
		queryStringStruct.Set("page_size", fmt.Sprintf("%d", r.PageSize))
	}

	if r.Type > 0 {
		queryStringStruct.Set("type", fmt.Sprintf("%d", r.Type))
	}

	queryStr := queryStringStruct.Encode()
	resResponse := r.bRequest.HttpGet(queryStr)

	return resResponse
}
