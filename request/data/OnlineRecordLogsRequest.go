package data

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
	"net/url"
)

type OnlineRecordLogsRequest struct {
	bRequest  *request.BaseRequest
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_at"`
}

func NewOnlineRecordLogsRequest(roomId int64) *OnlineRecordLogsRequest {

	r := new(OnlineRecordLogsRequest)
	b := new(request.BaseRequest)

	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.DataOnlineRecordLogsRequestUri, roomId)

	return r
}

func (r *OnlineRecordLogsRequest) SetApiVersion(apiVersion string) *OnlineRecordLogsRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

/**
 * @Description: 请求响应
 * @receiver r
 * @param client
 * @return request.ApiResponse
 */
func (r *OnlineRecordLogsRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client

	queryStringStruct := url.Values{}
	queryStringStruct.Set("start_at", fmt.Sprintf("%s", r.StartDate))
	queryStringStruct.Set("end_at", fmt.Sprintf("%s", r.EndDate))
	resResponse := r.bRequest.HttpGet(queryStringStruct.Encode())

	return resResponse
}
