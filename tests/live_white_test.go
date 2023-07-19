package tests

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request/live_white"
	"testing"
)

func TestLiveWhiteCreate(t *testing.T) {
	c := getClient()
	n := live_white.NewWhiteCreateRequest(469419621)
	n.Type = 1
	n.Content = "17845678069"
	r := n.Request(c)
	fmt.Println(r)
}

func TestLiveWhiteDelete(t *testing.T) {
	c := getClient()
	n := live_white.NewWhiteDeleteRequest(469419621)
	n.Type = 1
	n.Content = "17845678069"
	r := n.Request(c)
	fmt.Println(r)
}

func TestLiveWhiteList(t *testing.T) {
	c := getClient()
	n := live_white.NewWhiteListRequest(469419621)
	r := n.Request(c)
	fmt.Println(r)
}