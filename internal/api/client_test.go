package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Endpoints(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/3/movie/now_playing", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-token" {
			http.Error(w, `{"status_code":7,"status_message":"Invalid API key","success":false}`, http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"page":1,"results":[{"id":101,"title":"Now Playing Movie","vote_average":7.8,"vote_count":120}]}`))
	})
	mux.HandleFunc("/3/movie/popular", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"page":1,"results":[{"id":102,"title":"Popular Movie","vote_average":8.2,"vote_count":450}]}`))
	})
	mux.HandleFunc("/3/movie/top_rated", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"page":1,"results":[{"id":103,"title":"Top Rated Movie","vote_average":9.1,"vote_count":900}]}`))
	})
	mux.HandleFunc("/3/movie/upcoming", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"page":1,"results":[{"id":104,"title":"Upcoming Movie","vote_average":6.9,"vote_count":50}]}`))
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := NewClient("test-token")
	// Swap http client transport to redirect baseURL to httptest server
	client.httpClient = &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			req.URL.Scheme = "http"
			req.URL.Host = server.Listener.Addr().String()
			return http.DefaultTransport.RoundTrip(req)
		}),
	}

	ctx := context.Background()

	// Now Playing
	np, err := client.GetNowPlaying(ctx)
	if err != nil {
		t.Fatalf("unexpected error on GetNowPlaying: %v", err)
	}
	if len(np.Results) != 1 || np.Results[0].Title != "Now Playing Movie" {
		t.Errorf("unexpected results for GetNowPlaying: %+v", np.Results)
	}

	// Popular
	pop, err := client.GetPopular(ctx)
	if err != nil {
		t.Fatalf("unexpected error on GetPopular: %v", err)
	}
	if len(pop.Results) != 1 || pop.Results[0].Title != "Popular Movie" {
		t.Errorf("unexpected results for GetPopular: %+v", pop.Results)
	}

	// Top Rated
	tr, err := client.GetTopRated(ctx)
	if err != nil {
		t.Fatalf("unexpected error on GetTopRated: %v", err)
	}
	if len(tr.Results) != 1 || tr.Results[0].Title != "Top Rated Movie" {
		t.Errorf("unexpected results for GetTopRated: %+v", tr.Results)
	}

	// Upcoming
	up, err := client.GetUpcoming(ctx)
	if err != nil {
		t.Fatalf("unexpected error on GetUpcoming: %v", err)
	}
	if len(up.Results) != 1 || up.Results[0].Title != "Upcoming Movie" {
		t.Errorf("unexpected results for GetUpcoming: %+v", up.Results)
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
