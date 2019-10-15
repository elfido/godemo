package channel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var(
	ChannelURL = "http://localhost:8080/channel"
)

// Channel network channel information
type Channel struct {
	ID string `json:"ID"`
	Name string `json:"name"`
	Number int `json:"number"`
	Markets []string `json:"markets"`
}

// GetChannelInfo returns the data of a given channel from a remote service
func GetChannelInfo(channelID string) (Channel, error) {
	fullURL := fmt.Sprintf("%s/%s", ChannelURL, channelID)
	res, err := http.Get(fullURL)
	if err != nil {
		return Channel{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Channel{}, fmt.Errorf("invalid status code %d", res.StatusCode)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Channel{}, err
	}
	var channel Channel
	err = json.Unmarshal(bytes, &channel)
	if err != nil {
		return Channel{}, err
	}
	return channel, nil
}
