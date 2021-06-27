package tests

import (
	"fmt"
	"github.com/yuzihui/go-meetig-sdk/request/live_fly"
	"testing"
)

func TestLiveFlyCheckList(t *testing.T) {

	client := getClient()
	liveRequest := live_fly.NewLiveFlyCheckRequest()

	res := liveRequest.Request(client)
	fmt.Println(res)
}

func TestLiveFlyMemberList(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	liveRequest := live_fly.NewLiveFlyMemberRequest(634387860)

	res := liveRequest.Request(client)
	fmt.Println(res)
}
