package tests

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request/live"
	"testing"
)

func TestLiveList(t *testing.T) {

	client := getClient()
	liveRequest := live.NewLiveListRequest()
	liveRequest.Page = 1
	liveRequest.PageSize = 15

	res := liveRequest.Request(client)
	fmt.Println(res)

}

func TestLiveCreate(t *testing.T) {

	client := getClient()
	liveRequest := live.NewCreateRequest()
	liveRequest.Title = "go test meeting 19"
	liveRequest.StartAt = "2021-06-19 19:00:00"
	liveRequest.EndAt = "2021-06-19 22:00:00"
	liveRequest.UserId = 470
	liveRequest.LiveType = 1
	liveRequest.Introduction = "这是描述"
	liveRequest.IsSpeakerOnePassword = 1

	speakers := []map[string]interface{}{
		{"speaker_user_id": 456, "speaker_password": "test123"},
		{"speaker_user_id": 457, "speaker_password": "test123"},
	}
	liveRequest.Speakers = speakers
	liveRequest.AccessType = 1
	liveRequest.WatchPassword = "123456"
	res := liveRequest.SetApiVersion("v2").Request(client)

	fmt.Println(res)

}

func TestLiveUpdate(t *testing.T) {

	//3044073477
	client := getClient().SetHttpDebug(true)
	liveRequest := live.NewUpdateRequest(634387860)
	liveRequest.Title = "go test meeting1-test"
	liveRequest.StartAt = "2021-06-10 23:00:00"
	liveRequest.EndAt = "2021-06-10 23:30:00"
	liveRequest.UserId = 470
	liveRequest.LiveType = 1
	liveRequest.Introduction = "这是描述"
	liveRequest.IsSpeakerOnePassword = 0

	//speakers := []map[string]interface{}{
	//	{"speaker_user_id": 456, "speaker_password": "test123"},
	//	{"speaker_user_id": 457, "speaker_password": "test123"},
	//}
	//liveRequest.Speakers = speakers
	liveRequest.AccessType = 1
	liveRequest.WatchPassword = "123456"
	res := liveRequest.SetApiVersion("v2").Request(client)

	fmt.Println(res)

}

func TestNewLiveDetailRequest(t *testing.T) {
	client := getClient()
	liveRequest := live.NewLiveDetailRequest(3044073477)

	res := liveRequest.Request(client)

	fmt.Println(res)

}

func TestNewLiveDeleteRequest(t *testing.T) {
	client := getClient()

	liveRequest := live.NewLiveDeleteRequest(3044073477)

	res := liveRequest.Request(client)
	fmt.Println(res)

}

func TestNewLiveStatusRequest(t *testing.T) {
	client := getClient()

	liveRequest := live.NewLiveStatusRequest(839861479)

	liveRequest.Status = 1

	res := liveRequest.Request(client)
	fmt.Println(res)

}

func TestNewLiveAccessRequest(t *testing.T) {
	client := getClient()

	liveRequest := live.NewLiveAccessRequest(634387860)

	liveRequest.UserId = 458
	liveRequest.RoomPassword = "123456"
	liveRequest.RoomRole = 3

	res := liveRequest.Request(client)
	fmt.Println(res)

}

func TestNewLiveAddSpeakerRequest(t *testing.T) {
	client := getClient()

	liveRequest := live.NewLiveAddSpeakerRequest(634387860)

	speakers := []map[string]interface{}{
		{"speaker_user_id": 456, "speaker_password": "test123"},
		{"speaker_user_id": 457, "speaker_password": "test123"},
	}

	liveRequest.Speakers = speakers

	res := liveRequest.Request(client)
	fmt.Println(res)

}

func TestNewLiveRecordRequest(t *testing.T) {
	client := getClient().SetHttpDebug(true)

	liveRequest := live.NewLiveRecordRequest(634387860)

	liveRequest.All = 0

	res := liveRequest.Request(client)
	fmt.Println(res)

}

func TestNewLiveScreenShotRequest(t *testing.T) {
	client := getClient().SetHttpDebug(true)

	liveRequest := live.NewLiveScreenShotRequest(634387860)

	res := liveRequest.Request(client)
	fmt.Println(res)

}
