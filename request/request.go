package request

import (
	"encoding/json"
	"fmt"
	"github.com/yuzihui/go-meetig-sdk"
	"github.com/yuzihui/go-meetig-sdk/tools"
	"time"
)

type BaseRequest struct {
	Client        *meeting.MeetClient
	ApiVersion    string
	ApiRequestUri string
	Method        string
}

/**
 * @Description: 响应体
 */
type ApiResponse struct {
	Code    uint32      `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

/**
 * @Description: set header params
 * @receiver br
 */
func (br *BaseRequest) setHeader() {

	appId := br.Client.GetAppId()
	appSecret := br.Client.GetAppSecret()
	timestamp := time.Now().Unix()

	signature := tools.GetOpenApiSignature(appId, appSecret, timestamp)

	br.Client.TlkRequest.SetHeaders(map[string]string{
		"appid":     appId,
		"signature": signature,
		"timestamp": fmt.Sprintf("%d", timestamp),
	})
}

func (br *BaseRequest) getPathUri() string {
	return fmt.Sprintf("%s/%s", br.ApiVersion, br.ApiRequestUri)
}

/**
 * @Description: 获取数据json 响应
 * @param responseStr
 * @return ApiResponse
 */
func getResResponse(responseStr string) ApiResponse {
	var res ApiResponse
	json.Unmarshal([]byte(responseStr), &res)
	return res
}

/**
 * @Description: GET 请求
 * @receiver br
 * @param queryString
 * @return string
 */
func (br *BaseRequest) HttpGet(queryString string) ApiResponse {
	br.setHeader()
	path := br.getPathUri()

	requestClient := br.Client.TlkRequest.R()

	if len(queryString) > 0 {
		requestClient = requestClient.SetQueryString(queryString)
	}
	response, err := requestClient.Get(path)

	if err != nil {
		panic(err.Error())
	}
	return getResResponse(response.String())
}

/**
 * @Description: POST 请求
 * @receiver br
 * @param body []byte
 * @return string
 */
func (br *BaseRequest) HttpUpload(files map[string]string, body []byte) ApiResponse {
	br.setHeader()
	path := br.getPathUri()

	requestClient := br.Client.TlkRequest.R()

	if len(files) > 0 {
		requestClient = requestClient.SetFiles(files)
	}

	if len(body) > 0 {
		requestClient = requestClient.SetBody(body)
	}
	response, err := requestClient.Post(path)
	if err != nil {
		panic(err.Error())
	}
	return getResResponse(response.String())
}

/**
 * @Description: POST 请求
 * @receiver br
 * @param body []byte
 * @return string
 */
func (br *BaseRequest) HttpPost(body []byte) ApiResponse {
	br.setHeader()
	path := br.getPathUri()

	requestClient := br.Client.TlkRequest.R().SetHeader("Content-Type", "application/json")
	if len(body) > 0 {
		requestClient = requestClient.SetBody(body)
	}
	response, err := requestClient.Post(path)
	if err != nil {
		panic(err.Error())
	}
	return getResResponse(response.String())
}

/**
 * @Description: PUT 请求
 * @receiver br
 * @param map[string]interface{}
 * @return string
 */
func (br *BaseRequest) HttpPut(body []byte) ApiResponse {
	br.setHeader()
	path := br.getPathUri()

	requestClient := br.Client.TlkRequest.R().SetHeader("Content-Type", "application/json")
	if len(body) > 0 {
		requestClient = requestClient.SetBody(body)
	}

	response, err := requestClient.Put(path)
	if err != nil {
		panic(err.Error())
	}
	return getResResponse(response.String())
}

/**
 * @Description: DELETE 请求
 * @receiver br
 * @param map[string]interface{}
 * @return string
 */
func (br *BaseRequest) HttpDelete(body []byte) ApiResponse {
	br.setHeader()
	path := br.getPathUri()

	requestClient := br.Client.TlkRequest.R().SetHeader("Content-Type", "application/json")

	if len(body) > 0 {
		requestClient = requestClient.SetBody(body)
	}
	response, err := requestClient.Delete(path)

	if err != nil {
		panic(err.Error())
	}

	return getResResponse(response.String())
}
