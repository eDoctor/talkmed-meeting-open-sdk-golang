package survey

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
	"net/url"
)

type SurveyListRequest struct {
	bRequest *request.BaseRequest
	Module   int64 `json:"module"`
	ModuleId int64 `json:"module_id"`
	IsShow   uint8 `json:"is_show,omitempty"`
}

/**
 * @Description:
 * @param roomId
 * @return *SurveyListRequest
 */
func NewSurveyListRequest() *SurveyListRequest {

	r := new(SurveyListRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = constants.SurveyListRequestUri
	r.IsShow = 2 //默认值
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *SurveyListRequest
 */
func (r *SurveyListRequest) SetApiVersion(apiVersion string) *SurveyListRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *SurveyListRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client

	queryStringStruct := url.Values{}
	queryStringStruct.Set("is_show", fmt.Sprintf("%d", r.IsShow))

	if r.ModuleId > 0 {
		queryStringStruct.Set("module_id", fmt.Sprintf("%d", r.ModuleId))
	}

	if r.Module > 0 {
		queryStringStruct.Set("module", fmt.Sprintf("%d", r.Module))
	}

	queryStr := queryStringStruct.Encode()
	resResponse := r.bRequest.HttpGet(queryStr)

	return resResponse
}
