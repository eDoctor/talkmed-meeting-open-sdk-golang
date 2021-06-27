package tests

import (
	"github.com/yuzihui/go-meetig-sdk/request/survey"
	"testing"
)

func TestSurveyCreateRequest(t *testing.T) {

	client := getClient()
	survey := survey.NewSurveyCreateRequest()
	survey.Title = "调研测试2"
	survey.Introduction = "调研描述 go sdk"
	survey.Module = 1
	survey.ModuleId = 143427591
	survey.IsAnswerShow = 1
	res := survey.Request(client)

	t.Log(res)
}

func TestSurveyUpdateRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	survey := survey.NewSurveyUpdateRequest(280)
	survey.Title = "调研测试1go "
	survey.Introduction = "调研描述go sdk"
	//survey.IsShow = 0
	survey.IsAnswerShow = 0
	res := survey.Request(client)

	t.Log(res)

}

func TestSurveyDeleteRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	survey := survey.NewSurveyDeleteRequest(280)

	res := survey.Request(client)
	t.Log(res)

}

func TestSurveyListRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	survey := survey.NewSurveyListRequest()
	survey.ModuleId = 143427591
	survey.Module = 1
	survey.IsShow = 1

	res := survey.Request(client)
	t.Log(res)

}

func TestSurveyDetailRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	survey := survey.NewSurveyDetailRequest(281)
	res := survey.Request(client)
	t.Log(res)

}
