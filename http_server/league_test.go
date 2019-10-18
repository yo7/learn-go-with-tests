package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 30},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := StubPlayerStore{league: wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response)

		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, "application/json")
	})
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := response.Result().Header.Get("content-type")
	if got != want {
		t.Errorf("response did not have content-type of %q, got %q", want, got)
	}
}

func getLeagueFromResponse(t *testing.T, response *httptest.ResponseRecorder) (league []Player) {
	t.Helper()

	err := json.NewDecoder(response.Body).Decode(&league)

	if err != nil {
		t.Errorf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, league)
	}

	return
}

func assertLeague(t *testing.T, got, wanted []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("got %v want %v", got, wanted)
	}
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}
