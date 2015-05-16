package meerkat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	meerkatEndpoint = "https://%s.meerkatapp.co/%s"
)

const (
	version      = "1.0"
	apiURI       = "api"
	resourcesURI = "resources"
	channelsURI  = "channels"
)

const (
	routeBroadcastActivities   = "broadcastActivities"
	routeLiveNow               = "liveNow"
	routeBroadcastRestreams    = "broadcastRestreams"
	routeBroadcastWatchers     = "broadcastWatchers"
	routeProfile               = "profile"
	routeScheduledStreams      = "scheduledStreams"
	routeBroadcastLikes        = "broadcastLikes"
	routeStreamSummaryTemplate = "streamSummaryTemplate"
	routeBroadcastComments     = "broadcastComments"
)

// Client represents an Meerkat API client.
type Client struct {
	token  string
	routes map[string]string
}

// NewClient returns a client.
func NewClient(token string) (*Client, error) {
	c := &Client{
		token: token,
	}
	err := c.setRoutes()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// setRoutes requests the route map from Meerkat
// and sets the map on the client.
func (c *Client) setRoutes() error {
	res, err := c.get(fmt.Sprintf(meerkatEndpoint, apiURI, "routes"))
	if err != nil {
		return err
	}

	var resBody map[string]string
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		return err
	}

	c.routes = resBody
	return nil
}

// get conducts a GET request on the meerkat api. It sets
// the version in the url and sets the Authorization header.
func (c *Client) get(url string) (*http.Response, error) {
	url += "?v=" + version
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.token)
	return http.DefaultClient.Do(req)
}

// GetAllBroadcasts gets all current broadcasts.
func (c *Client) GetAllBroadcasts() ([]*Broadcast, error) {
	url := c.routes[routeLiveNow]
	res, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var resBody struct {
		Result []*Broadcast
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return resBody.Result, nil
}

// GetBroadcast gets a summary of a broadcast.
func (c *Client) GetBroadcast(id string) (*Broadcast, error) {
	url := strings.Replace(c.routes[routeStreamSummaryTemplate], "{broadcastId}", id, -1)
	res, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var resBody struct {
		Result *Broadcast
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return resBody.Result, nil
}

// GetBroadcastActivities gets a broadcast and it's associated activities.
func (c *Client) GetBroadcastActivities(id string) (*Broadcast, error) {
	url := strings.Replace(c.routes[routeBroadcastActivities], "{broadcastId}", id, -1)
	res, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var resBody struct {
		Result *Broadcast
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return resBody.Result, nil
}

// GetBroadcastRestreams gets restreams for a broadcast.
func (c *Client) GetBroadcastRestreams(id string) ([]*Activity, error) {
	url := strings.Replace(c.routes[routeBroadcastRestreams], "{broadcastId}", id, -1)
	res, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var resBody struct {
		Result []*Activity
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return resBody.Result, nil
}

// GetBroadcastComments gets comments for a broadcast.
func (c *Client) GetBroadcastComments(id string) ([]*Activity, error) {
	url := strings.Replace(c.routes[routeBroadcastComments], "{broadcastId}", id, -1)
	res, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var resBody struct {
		Result []*Activity
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return resBody.Result, nil
}

// GetBroadcastLikes gets likes for a broadcast.
func (c *Client) GetBroadcastLikes(id string) ([]*Activity, error) {
	url := strings.Replace(c.routes[routeBroadcastLikes], "{broadcastId}", id, -1)
	res, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var resBody struct {
		Result []*Activity
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return resBody.Result, nil
}

// GetScheduledBroadcasts gets upcoming broadcasts.
func (c *Client) GetScheduledBroadcasts() ([]*Broadcast, error) {
	url := c.routes[routeScheduledStreams]
	res, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var resBody struct {
		Result []*Broadcast
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return resBody.Result, nil
}

// GetUserDetails gets user info.
func (c *Client) GetUserDetails(id string) (*User, error) {
	url := strings.Replace(c.routes[routeProfile], "{userId}", id, -1)
	res, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var resBody struct {
		Result *User
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return resBody.Result, nil
}
