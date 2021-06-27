package tests

import (
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request/menu"
	"testing"
)

/**
 * @Description:
 * @param t
 */
func TestMenuCreateRequest(t *testing.T) {
	client := getClient()
	menu := menu.NewMenuCreateRequest(143427591)
	menu.MenuTitle = "测试2"
	menu.MenuTitleEn = "test2"
	menu.MenuContent = "测试内容"
	menu.MenuShow = 1
	menu.MenuType = 0
	menu.MenuHref = "www.baidu.com"

	res := menu.Request(client)

	t.Log(res)
}

/**
 * @Description:
 * @param t
 */
func TestMenuUpdateRequest(t *testing.T) {
	client := getClient()
	menu := menu.NewMenuUpdateRequest(143427591, "53")
	menu.MenuTitle = "测试3"
	menu.MenuTitleEn = "test3"
	menu.MenuContent = "测试内容"
	menu.MenuShow = 1
	menu.MenuType = 0
	menu.MenuHref = "www.baidu.com"

	res := menu.Request(client)
	t.Log(res)
}

/**
 * @Description:
 * @param t
 */
func TestMenuListRequest(t *testing.T) {
	client := getClient()
	menu := menu.NewMenuListRequest(143427591)

	res := menu.Request(client)
	t.Log(res)
}

/**
 * @Description:
 * @param t
 */
func TestMenuDeleteRequest(t *testing.T) {
	client := getClient()
	menu := menu.NewMenuDeleteRequest(143427591, "54")

	res := menu.Request(client)
	t.Log(res)
}

/**
 * @Description:
 * @param t
 */
func TestMenuSortRequest(t *testing.T) {
	client := getClient().SetHttpDebug(true)
	menu := menu.NewMenuSortRequest(143427591)
	menu.Sort = []string{"53", "survey", "chats", "files"}

	res := menu.Request(client)
	t.Log(res)
}
