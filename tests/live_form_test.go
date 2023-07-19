package tests

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request/live_form"
	"testing"
)

func TestFormBind(t *testing.T) {
	c := getClient()
	n := live_form.NewLiveBindFormRequest()
	n.FormId = 37
	n.RoomId = 2588879831
	n.CollectionRule = 1
	r := n.Request(c)
	fmt.Println(r)
}
