package survey

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request"
	constants "github.com/eDoctor/talkmed-meeting-open-sdk-golang/tools/constant"
)

/**
 * @Description:添加调研请求
 */
type DeleteRequest struct {
	bRequest *request.BaseRequest
}

/**
 * @Description:
 *
 */
func NewSurveyDeleteRequest(surveyId int64) *DeleteRequest {

	r := new(DeleteRequest)
	b := new(request.BaseRequest)
	r.bRequest = b
	r.bRequest.ApiVersion = constants.ApiV1
	r.bRequest.ApiRequestUri = fmt.Sprintf(constants.SurveyDeleteRequestUri, surveyId)
	return r
}

/**
 * @Description:
 * @receiver r
 * @param apiVersion
 *
 */
func (r *DeleteRequest) SetApiVersion(apiVersion string) *DeleteRequest {
	r.bRequest.ApiVersion = apiVersion
	return r
}

func (r *DeleteRequest) Request(client *meeting.MeetClient) request.ApiResponse {

	r.bRequest.Client = client
	return r.bRequest.HttpDelete([]byte(""))
}
