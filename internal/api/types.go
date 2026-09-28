package api

import "fmt"

// Movie represents a single movie returned by the TMDB API.
type Movie struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	OriginalLanguage string  `json:"original_language"`
	Popularity       float64 `json:"popularity"`
}

// MovieListResponse represents the paginated response for a list of movies.
type MovieListResponse struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

// Error represents an error returned by TMDB.
type Error struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("tmdb api error (code %d): %s", e.StatusCode, e.StatusMessage)
}
