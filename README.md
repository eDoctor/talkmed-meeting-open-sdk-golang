### TalkMed Meeting Open Api Sdk

### 其他版本

- [php sdk](https://github.com/eDoctor/talkmed-meeting-open-sdk)

### Wiki地址

```
https://wiki.meeting.talkmed.com
```

### 安装

```
go get github.com/eDoctor/talkmed-meeting-open-sdk-golang
```

### 测试用例

```

func TestLiveCheckRequest(t *testing.T) {
			client := meeting.NewMeetingClint(AppId, AppSecret).SetHost(DevBaseUri).SetHttpDebug(false)		
      liveRequest := live_chat.NewLiveChatCheckRequest(2089596951, 986)
      liveRequest.AuditStatus = 2
      res := liveRequest.Request(client)
      t.log(res)
}    
```

###客户端

```
client := meeting.NewMeetingClint(AppId, AppSecret).SetHost(DevBaseUri).SetHttpDebug(false)		
```

方法说明:

```
SetHost 设置域名 例:devapimeeting.talkmed.com
SetHttpDebug http请求打印输出
```

### 请求入参

```
  liveRequest := live_chat.NewLiveChatCheckRequest(2089596951, 986)
  liveRequest.AuditStatus = 2
  res := liveRequest.SetApiVersion('v1').Request(client)
```

说明：

```
请求的结构体实例化都是New开头
NewLiveChatCheckRequest  //请求参数对象构造函数会初始化请求uri
SetApiVersion //该方法可设置请求的接口版本 默认v1
具体参数参照wiki与ide提示
```

### 模块

```
用户相关接口: user
会议相关接口: live
会议飞检相关接口: live_fly
会议讨论相关接口: live_chat
文件相关接口: live_file
调研相关: survey
调研问题相关: survey_question
菜单相关: survey_question
会后数据报表接口: data
```

说明

```
会后数据报表接口(data)内的方法，可根据接口地址的最后的拼接单词获取new实例，例如:
		v1/open/room/{roomId:[0-9]+}/live_base   live_base在sdk中对应的是 NewLiveBaseRequest
		v1/open/room/{roomId:[0-9]+}/user_join_logs user_join_logs在sdk中对应的是 NewUserJoinLogsRequest
非data模块接口,需根据接口提示的增删改查与接口地址自行调用
```

### 参会地址跳转

```
func TestGenerateJoinUrl(t *testing.T) {
	str := tools.GetAuthorizeUri("https://devmeeting.talkmed.com", "xxxx", "xxxx", "e96a4dba-4eb2-dd1d-7fa3-bdfcd36d98d7", "1631372727", "2", "web", "", "")
	t.Log(str)
}
```

例子：

```
https://devmeeting.talkmed.com/oauth/authorize?app_id=xxxx&auth_token=e96a4dba-4eb2-dd1d-7fa3-bdfcd36d98d7&channel=&password=&platform=web&role=2&room_id=1631372727&signature=8c5a485b5ced25231f0f290adb312ac62f8547c4a08d339427a821e12a1335f0&timestamp=1624771638
```

