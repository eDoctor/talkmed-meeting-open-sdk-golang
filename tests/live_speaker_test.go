package tests

import (
	"fmt"
	"github.com/eDoctor/talkmed-meeting-open-sdk-golang/request/live_speaker"
	"testing"
)

func TestLiveSpeakerCreate(t *testing.T) {
	c := getClient()
	n := live_speaker.NewSpeakerCreateRequest(469419621)
	s := []map[string]interface{}{
		{"speaker_user_id": 45030, "speaker_password": "test123"},
	}
	n.Speakers = s
	r := n.Request(c)
	fmt.Println(r)
}

func TestLiveSpeakerUpdate(t *testing.T) {
	c := getClient()
	n := live_speaker.NewSpeakerCreateRequest(469419621)
	s := []map[string]interface{}{
		{"speaker_user_id": 45030, "speaker_password": "test123"},
	}
	n.Speakers = s
	r := n.Request(c)
	fmt.Println(r)
}

func TestLiveSpeakerDelete(t *testing.T) {
	c := getClient()
	n := live_speaker.NewSpeakerCreateRequest(469419621)
	s := []map[string]interface{}{
		{"speaker_user_id": 45030, "speaker_password": "test123"},
	}
	n.Speakers = s
	r := n.Request(c)
	fmt.Println(r)
}

func TestLiveSpeakerList(t *testing.T) {
	c := getClient()
	n := live_speaker.NewSpeakerListRequest(469419621)
	r := n.Request(c)
	fmt.Println(r)
}

func TestLiveSpeakerDetail(t *testing.T) {
	c := getClient()
	n := live_speaker.NewSpeakerCreateRequest(469419621)
	s := []map[string]interface{}{
		{"speaker_user_id": 45030, "speaker_password": "test123"},
	}
	n.Speakers = s
	r := n.Request(c)
	fmt.Println(r)
}