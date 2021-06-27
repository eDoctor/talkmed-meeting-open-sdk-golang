package menu

import (
	"encoding/json"
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/request"
	constants "github.com/yuzihui/go-meetig-sdk/tools/constant"
)

/**
 * @Description:状态请求
 */
type SortRequest struct {
	bRequest *request.BaseRequest
	Sort     []string `json:"sort"`
}

/**
 * @Description:
 *
 */
func NewMenuSortRequest(roomId int64) *SortRequest {

	r := new(SortRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.MenuSortRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 *
 */
func (r *SortRequest) SetApiVersion(apiVersion string) *SortRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *SortRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	paramsByte, _ := json.Marshal(r)

	return r.bRequest.HttpPost(paramsByte)
}
