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
type CreateRequest struct {
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
 * @return *CreateRequest
 */
func NewMenuCreateRequest(roomId int64) *CreateRequest {

	r := new(CreateRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.MenuCreateRequestUri, roomId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 * @return *CreateRequest
 */
func (r *CreateRequest) SetApiVersion(apiVersion string) *CreateRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *CreateRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	registerByte, _ := json.Marshal(r)

	fmt.Println(string(registerByte))
	return r.bRequest.HttpPost(registerByte)
}
