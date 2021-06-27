package tools

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

/**
 * 获取open api 签名
 */
func GetOpenApiSignature(appId string, appSecret string, timestamp int64) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s-%s-%d", appId, appSecret, timestamp)))
	r := h.Sum(nil)
	signature := fmt.Sprintf("%x", r)
	return signature
}

/**
 * 获取授权Uri
 */
func GetAuthorizeUri(host string, appId string, appSecret string, authToken string, roomId string, role string, platform string, password string, channel string) string {

	timestamp := time.Now().Unix()
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s-%s-%s-%d", appId, appSecret, authToken, timestamp)))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	var uri url.URL
	query := uri.Query()
	query.Set("app_id", appId)
	query.Set("auth_token", authToken)
	query.Set("timestamp", strconv.Itoa(int(timestamp)))
	query.Set("signature", signature)
	query.Set("platform", platform)
	query.Set("room_id", roomId)
	query.Set("role", role)
	query.Set("channel", channel)
	query.Set("password", password)

	return fmt.Sprintf("%s/oauth/authorize?%s", host, query.Encode())
}

func StructToMapByJsonTag(i interface{}) map[string]interface{} {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	data := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		filed := f.Tag.Get("json")
		if len(filed) == 0 {
			continue
		}
		val := v.Field(i).Interface()
		data[filed] = val
	}
	return data
}

func HttpBuildQuery(params map[string]interface{}) string {

	Url := url.Values{}
	for k, v := range params {
		Url.Set(k, v.(string))
	}
	return Url.Encode()
}
