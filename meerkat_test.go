package meerkat

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unrolled/render"
)

var (
	testClient *Client
	renderer   = render.New(render.Options{
		IndentJSON:    true,
		IsDevelopment: true,
	})
	testBroadcastId = "test-broadcast-id"
	testUserId      = "test-user-id"
)

func init() {
	testClient = &Client{
		token:  "test-token",
		routes: map[string]string{},
	}
}

func TestGetAllBroadcasts(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(testGetAllBroadcastsHandler))
	defer server.Close()
	testClient.routes[routeLiveNow] = server.URL

	_, err := testClient.GetAllBroadcasts()
	if err != nil {
		t.Error(err)
	}
}

func testGetAllBroadcastsHandler(rw http.ResponseWriter, req *http.Request) {
	renderer.JSON(rw, http.StatusOK, map[string]interface{}{
		"result": []Broadcast{},
	})
}

func TestGetBroadcast(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(testGetBroadcastHandler))
	defer server.Close()
	testClient.routes[routeStreamSummaryTemplate] = server.URL

	_, err := testClient.GetBroadcast(testBroadcastId)
	if err != nil {
		t.Error(err)
	}
}

func testGetBroadcastHandler(rw http.ResponseWriter, req *http.Request) {
	renderer.JSON(rw, http.StatusOK, map[string]interface{}{
		"result": Broadcast{},
	})
}

func TestGetBroadcastActivities(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(testGetBroadcastActivitiesHandler))
	defer server.Close()
	testClient.routes[routeBroadcastActivities] = server.URL

	_, err := testClient.GetBroadcastActivities(testBroadcastId)
	if err != nil {
		t.Error(err)
	}
}

func testGetBroadcastActivitiesHandler(rw http.ResponseWriter, req *http.Request) {
	renderer.JSON(rw, http.StatusOK, map[string]interface{}{
		"result": Broadcast{},
	})
}

func TestGetBroadcastRestreams(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(testGetBroadcastRestreamsHandler))
	defer server.Close()
	testClient.routes[routeBroadcastRestreams] = server.URL

	_, err := testClient.GetBroadcastRestreams(testBroadcastId)
	if err != nil {
		t.Error(err)
	}
}

func testGetBroadcastRestreamsHandler(rw http.ResponseWriter, req *http.Request) {
	renderer.JSON(rw, http.StatusOK, map[string]interface{}{
		"result": []Activity{},
	})
}

func TestGetBroadcastComments(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(testGetBroadcastCommentsHandler))
	defer server.Close()
	testClient.routes[routeBroadcastComments] = server.URL

	_, err := testClient.GetBroadcastComments(testBroadcastId)
	if err != nil {
		t.Error(err)
	}
}

func testGetBroadcastCommentsHandler(rw http.ResponseWriter, req *http.Request) {
	renderer.JSON(rw, http.StatusOK, map[string]interface{}{
		"result": []Activity{},
	})
}

func TestGetBroadcastLikes(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(testGetBroadcastLikesHandler))
	defer server.Close()
	testClient.routes[routeBroadcastLikes] = server.URL

	_, err := testClient.GetBroadcastLikes(testBroadcastId)
	if err != nil {
		t.Error(err)
	}
}

func testGetBroadcastLikesHandler(rw http.ResponseWriter, req *http.Request) {
	renderer.JSON(rw, http.StatusOK, map[string]interface{}{
		"result": []Activity{},
	})
}

func TestGetScheduledBroadcasts(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(testGetScheduledBroadcastsHandler))
	defer server.Close()
	testClient.routes[routeScheduledStreams] = server.URL

	_, err := testClient.GetScheduledBroadcasts()
	if err != nil {
		t.Error(err)
	}
}

func testGetScheduledBroadcastsHandler(rw http.ResponseWriter, req *http.Request) {
	renderer.JSON(rw, http.StatusOK, map[string]interface{}{
		"result": []Broadcast{},
	})
}

func TestGetUserDetails(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(testGetUserDetailsHandler))
	defer server.Close()
	testClient.routes[routeProfile] = server.URL

	_, err := testClient.GetUserDetails(testUserId)
	if err != nil {
		t.Error(err)
	}
}

func testGetUserDetailsHandler(rw http.ResponseWriter, req *http.Request) {
	renderer.JSON(rw, http.StatusOK, map[string]interface{}{
		"result": User{},
	})
}
