package channel

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetChannelInfoBadPayload(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte{1, 2, 3})
	}))

	ChannelURL = ts.URL
	channel, err := GetChannelInfo("123")
	if channel.ID != "" {
		t.Errorf("Expected channel to be empty, instead found channeld ID = %s", channel.ID)
	}

	if err == nil {
		t.Error("Expected error to be related to invalid payload, instead found nil")
	}
}

func TestGetChannelInfo(t *testing.T) {
	const channelName = "The xzy channel"
	const channelID = "123"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		channel := Channel{
			ID:      channelID,
			Name:    channelName,
			Number:  231,
			Markets: []string{},
		}
		bytes, _ := json.Marshal(channel)
		w.Write(bytes)
	}))

	ChannelURL = ts.URL
	channel, err := GetChannelInfo("123")
	if channel.ID != channelID {
		t.Errorf("Expected channel %s, instead found channeld ID = %s", channelID, channel.ID)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, instead found %s", err)
	}
}

func TestGetChannelInfoBadStatusCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}))

	ChannelURL = ts.URL
	channel, err := GetChannelInfo("123")
	if channel.ID != "" {
		t.Errorf("Expected an empty channel, instead found channeld ID = %s", channel.ID)
	}
	if err == nil {
		t.Error("Expected error to report bad status code, instead found nil error")
	}
}