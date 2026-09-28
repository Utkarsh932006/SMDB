package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.themoviedb.org/3"

// Client handles communication with the TMDB API.
type Client struct {
	token      string
	httpClient *http.Client
}

// NewClient creates a new TMDB API client.
func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: http.DefaultClient,
	}
}

// GetNowPlaying fetches the now playing movies.
func (c *Client) GetNowPlaying(ctx context.Context) (*MovieListResponse, error) {
	return c.fetch(ctx, "/movie/now_playing")
}

// GetPopular fetches popular movies.
func (c *Client) GetPopular(ctx context.Context) (*MovieListResponse, error) {
	return c.fetch(ctx, "/movie/popular")
}

// GetTopRated fetches top rated movies.
func (c *Client) GetTopRated(ctx context.Context) (*MovieListResponse, error) {
	return c.fetch(ctx, "/movie/top_rated")
}

// GetUpcoming fetches upcoming movies.
func (c *Client) GetUpcoming(ctx context.Context) (*MovieListResponse, error) {
	return c.fetch(ctx, "/movie/upcoming")
}

func (c *Client) fetch(ctx context.Context, endpoint string) (*MovieListResponse, error) {
	url := baseURL + endpoint
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request for %s: %w", endpoint, err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr Error
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err == nil {
			return nil, &apiErr
		}
		return nil, fmt.Errorf("unexpected status %d while fetching %s", resp.StatusCode, endpoint)
	}

	var result MovieListResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response for %s: %w", endpoint, err)
	}

	return &result, nil
}
