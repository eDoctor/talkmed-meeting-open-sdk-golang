package tests

import (
	"fmt"
	"github.com/eDoctor/meeting/request/live_file"
	"testing"
)

func TestLiveFileUploadRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	liveRequest := live_file.NewLiveFileUploadRequest()
	liveRequest.RoomId = 3645780814
	liveRequest.UserId = 470
	liveRequest.FilePath = "/data/list.pdf"

	res := liveRequest.Request(client)

	t.Log(res)
}

func TestLiveFileDeleteRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	liveRequest := live_file.NewLiveFileDeleteRequest(556)

	res := liveRequest.Request(client)

	fmt.Println(res)
}

func TestLiveFilePrintListRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	liveRequest := live_file.NewLiveFilePrintListRequest(682)
	liveRequest.UserId = 470
	res := liveRequest.Request(client)
	fmt.Println(res)
}

func TestLiveFileBindRequest(t *testing.T) {
	client := getClient()
	liveRequest := live_file.NewLiveFileBindRequest(3645780814)
	liveRequest.FileIds = []int64{561}
	res := liveRequest.Request(client)
	t.Log(res)
}

func TestLiveFileListBindRequest(t *testing.T) {
	client := getClient()
	liveRequest := live_file.NewLiveFileBindListRequest(3645780814)
	res := liveRequest.Request(client)
	t.Log(res)
}

func TestLiveFileBindDeleteRequest(t *testing.T) {
	client := getClient()
	liveRequest := live_file.NewLiveFileBindDeleteRequest(3645780814)
	res := liveRequest.Request(client)
	t.Log(res)
}
