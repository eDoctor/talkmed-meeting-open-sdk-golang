package tests

import (
	"github.com/yuzihui/go-meetig-sdk/request/data"
	"testing"
)

const DataRoomId = 285405984

func TestAboutLiveRequest(t *testing.T) {

	requestObj := data.NewAboutLiveRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)

	res := requestObj.Request(client)
	t.Log(res)

}

func TestAccessRecordsRequest(t *testing.T) {
	requestObj := data.NewAccessRecordsRequestRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)
	t.Log(res)

}

func TestLiveBaseRequest(t *testing.T) {
	requestObj := data.NewLiveBaseRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)
	t.Log(res)

}

func TestLiveReportRequest(t *testing.T) {
	requestObj := data.NewLiveReportRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)
	t.Log(res)

}

func TestOnlineRecordLogsRequest(t *testing.T) {
	requestObj := data.NewOnlineRecordLogsRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)
	t.Log(res)

}

func TestSurveyInfosRequest(t *testing.T) {
	requestObj := data.NewSurveyInfosRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)
	t.Log(res)

}

func TestSurveyRecordRequest(t *testing.T) {
	requestObj := data.NewSurveyRecordsRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestUserActionLogsRequest(t *testing.T) {
	requestObj := data.NewUserActionLogsRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestUserAllLogsRequest(t *testing.T) {
	requestObj := data.NewUserAllLogsRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestUserJoinLogsRequest(t *testing.T) {
	requestObj := data.NewUserJoinLogsRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestUserLiveAllLogsRequest(t *testing.T) {
	requestObj := data.NewUserLiveAllLogsRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestUserLiveLogsByRtcTimeRequest(t *testing.T) {
	requestObj := data.NewUserLiveLogsByRtcTimeRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestUserLiveRequest(t *testing.T) {
	requestObj := data.NewUserLiveRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestUserVodJoinLogsRequest(t *testing.T) {
	requestObj := data.NewUserVodJoinLogsRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestVodReportRequest(t *testing.T) {
	requestObj := data.NewVodRportRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestWebSocketChatsRequest(t *testing.T) {
	requestObj := data.NewWebsocketChatsRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}

func TestUserLiveUniqueLogsByRtcTimeRequest(t *testing.T) {
	requestObj := data.NewUserLiveUniqueLogsByRtcTimeRequest(DataRoomId)
	client := getClient().SetHttpDebug(true)
	res := requestObj.Request(client)

	t.Log(res)

}
