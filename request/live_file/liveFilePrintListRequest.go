package live_file

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
	"net/url"
)

type FilePrintListRequest struct {
	bRequest *request.BaseRequest
	UserId   int64
}

/**
 * @Description:
 * @param roomId
 * @return *FilePrintListRequest
 */
func NewLiveFilePrintListRequest(FielId int64) *FilePrintListRequest {

	r := new(FilePrintListRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.LiveFilePrintListRequestUri, FielId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *FilePrintListRequest
 */
func (r *FilePrintListRequest) SetApiVersion(apiVersion string) *FilePrintListRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *FilePrintListRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client

	queryStringStruct := url.Values{}
	if r.UserId > 0 {
		queryStringStruct.Set("user_id", fmt.Sprintf("%d", r.UserId))
	}

	queryStr := queryStringStruct.Encode()
	fmt.Println(queryStr)
	resResponse := r.bRequest.HttpGet(queryStr)

	return resResponse
}
