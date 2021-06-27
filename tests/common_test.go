package tests

import (
	"encoding/json"
	"fmt"
	"github.com/eDoctor/meeting"
	"github.com/eDoctor/meeting/tools"
	"github.com/go-resty/resty/v2"
	"testing"
	"time"
)

const AppId = "xxxxxxx"
const AppSecret = "xxxxxxxxxxxx"

const DevBaseUri = "https://devapimeeting.talkmed.com"

func getClient() *meeting.MeetClient {

	client := meeting.NewMeetingClint(AppId, AppSecret).SetHost(DevBaseUri).SetHttpDebug(false)
	return client
}

func TestClient(t *testing.T) {

	getClient()
	//var r request.Request

	//r.
	//
	//client, _ := meeting.NewMeetingClint(AppId, AppSecret).SetBaseUri(DevBaseUri).SetApiVersion("v1").Request(r)
	//
	//fmt.Println(client)

}

func TestRequest(t *testing.T) {

	timestamp := time.Now().Unix()
	signature := tools.GetOpenApiSignature(AppId, AppSecret, timestamp)

	client := resty.New()
	client.SetTimeout(1 * time.Minute)
	client.SetHostURL("https://devapimeeting.talkmed.com")
	client.SetHeaders(map[string]string{"appid": AppId, "signature": signature, "timestamp": fmt.Sprintf("%d", timestamp)})
	response, _ := client.R().Get("/v1/open/room")

	type Response struct {
		Code    uint        `json:"code"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	var res Response

	_ = json.Unmarshal([]byte(response.String()), &res)

	fmt.Println(res.Data)

}

func TestOpenApiSignature(t *testing.T) {
	signature := tools.GetOpenApiSignature(AppId, AppSecret, 1622963549)
	fmt.Println(signature)
}

func TestGenerateJoinUrl(t *testing.T) {

	str := tools.GetAuthorizeUri("https://devmeeting.talkmed.com", "xxxx", "xxxx", "e96a4dba-4eb2-dd1d-7fa3-bdfcd36d98d7", "1631372727", "2", "web", "", "")
	t.Log(str)
}
