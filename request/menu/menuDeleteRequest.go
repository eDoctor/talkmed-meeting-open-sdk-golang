package menu

import (
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/request"
	constants "github.com/eDoctor/meeting/tools/constant"
)

/**
 * @Description:添加调研请求
 */
type DeleteRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description:
 * @param roomId
 * @param menuId
 * @return *DeleteRequest
 */
func NewMenuDeleteRequest(roomId int64, menuSign string) *DeleteRequest {

	r := new(DeleteRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.MenuDeleteRequestUri, roomId, menuSign)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 */
func (r *DeleteRequest) SetApiVersion(apiVersion string) *DeleteRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *DeleteRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	return r.bRequest.HttpDelete([]byte(""))
}
