package meeting

import (
	"github.com/go-resty/resty/v2"
)

type MeetClient struct {
	TlkRequest *resty.Client
	AppId      string
	appSecret  string
}

/**
 * @Description:创建meeting 请求客户端
 * @param appId
 * @param appSecret
 * @return *MeetClient
 */
func NewMeetingClint(appId string, appSecret string) *MeetClient {
	meeting := &MeetClient{}
	meeting.AppId = appId
	meeting.appSecret = appSecret
	meeting.TlkRequest = resty.New()
	return meeting
}

/**
 * @Description: resty 报请求相关参数
 * @receiver m
 * @param isDebug
 * @return *MeetClient
 */
func (m *MeetClient) SetHttpDebug(isDebug bool) *MeetClient {
	m.TlkRequest.SetDebug(isDebug)
	return m
}

/**
 * @Description: 设置请求域名
 * @receiver m
 * @param host
 * @return *MeetClient
 */
func (m *MeetClient) SetHost(host string) *MeetClient {
	m.TlkRequest.SetHostURL(host)
	return m
}

func (m *MeetClient) GetAppId() string {
	return m.AppId
}

func (m *MeetClient) GetAppSecret() string {
	return m.appSecret
}
