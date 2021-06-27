package live

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
	"net/url"
)

type ListRequest struct {
	bRequest *request.BaseRequest
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
	Title    string `json:"title"`
}

func NewLiveListRequest() *ListRequest {

	r := new(ListRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.LiveListRequestUri

	return r
}

func (r *ListRequest) SetApiVersion(apiVersion string) *ListRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *ListRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client

	queryStringStruct := url.Values{}

	if r.Page > 0 {
		queryStringStruct.Set("page", fmt.Sprintf("%d", r.Page))
	}

	if r.PageSize > 0 {
		queryStringStruct.Set("page_size", fmt.Sprintf("%d", r.PageSize))
	}

	if len(r.Title) > 0 {
		queryStringStruct.Set("title", fmt.Sprintf("%s", r.Title))
	}
	queryStr := queryStringStruct.Encode()
	resResponse := r.bRequest.HttpGet(queryStr)

	return resResponse
}
