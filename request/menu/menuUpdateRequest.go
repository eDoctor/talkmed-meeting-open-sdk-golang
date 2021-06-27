package menu

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

/**
 * @Description:添加调研请求
 */
type UpdateRequest struct {
	bRequest    *request.BaseRequest
	MenuTitle   string `json:"menu_title"`
	MenuTitleEn string `json:"menu_title_en"`
	MenuType    uint8  `json:"menu_type"`
	MenuContent string `json:"menu_content,omitempty"`
	MenuHref    string `json:"module,omitempty"`
	MenuShow    uint8  `json:"menu_show"`
}

/**
 * @Description:
 * @param roomId
 * @param menuId
 * @return *UpdateRequest
 */
func NewMenuUpdateRequest(roomId int64, menuSign string) *UpdateRequest {

	r := new(UpdateRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.MenuUpdateRequestUri, roomId, menuSign)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 *
 */
func (r *UpdateRequest) SetApiVersion(apiVersion string) *UpdateRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *UpdateRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)
	return r.bRequest.HttpPut(registerByte)
}
