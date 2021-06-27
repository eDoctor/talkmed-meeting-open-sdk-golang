package tests

import (
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request/survey_question"
	"testing"
)

func TestSurveyQuestionCreateRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	survey := survey_question.NewQuestionCreateRequest(281)

	survey.Title = "今天的天气又如何555?"
	survey.QuestionType = 1
	survey.IsMust = 1

	//options := []survey_question.Options{{Title: "晴天", Introduction: "这是晴天"}, {Title: "阴天", Introduction: "这是阴天"}, {Title: "下雨", Introduction: "这是下雨"}, {Title: "下雪", Introduction: "这是下雪"}}

	//options := make([]survey_question.Options, 4)
	//options[0] = survey_question.Options{Title: "晴天", Introduction: "这是晴天"}
	//options[1] = survey_question.Options{Title: "阴天", Introduction: "这是阴天"}
	//options[2] = survey_question.Options{Title: "下雨", Introduction: "这是下雨"}
	//options[3] = survey_question.Options{Title: "下雪", Introduction: "这是下雪"}

	survey.Options = []map[string]interface{}{{"title": "晴天", "introduction": "这是晴天"}, {"title": "阴天", "introduction": "这是阴天"}}
	res := survey.Request(client)

	t.Log(res)
}

func TestSurveyQuestionUpdateRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	survey := survey_question.NewQuestionUpdateRequest(281, 451)
	survey.Title = "调研测试哈哈1111"
	survey.Introduction = "调研描述go sdk"

	survey.Options = []map[string]interface{}{{"id": 1293, "title": "晴天22222", "introduction": "这是晴天2222"}, {"id": 1294, "title": "阴天", "introduction": "这是阴天"}}

	res := survey.Request(client)

	t.Log(res)

}

func TestSurveyQuestionDeleteRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	survey := survey_question.NewQuestionDeleteRequest(281, 444)

	res := survey.Request(client)
	t.Log(res)

}

func TestSurveyQuestionListRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	survey := survey_question.NewQuestionListRequest(281)

	res := survey.Request(client)
	t.Log(res)

}

func TestSurveyQuestionDetailRequest(t *testing.T) {

	client := getClient().SetHttpDebug(true)
	survey := survey_question.NewQuestionDetailRequest(281, 451)
	res := survey.Request(client)
	t.Log(res)
}
