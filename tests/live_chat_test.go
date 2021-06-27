package tests

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk/request/live_chat"
	"testing"
)

func TestLiveCheckRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	liveRequest := live_chat.NewLiveChatCheckRequest(2089596951, 986)
	liveRequest.AuditStatus = 2

	res := liveRequest.Request(client)

	fmt.Println(res)
}

func TestLiveChatListRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	liveRequest := live_chat.NewLiveChatListRequest(2089596951)
	liveRequest.PageSize = 2
	liveRequest.Page = 2

	res := liveRequest.Request(client)
	fmt.Println(res)
}
