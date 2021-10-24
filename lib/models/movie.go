package models

import (
	contentsGrpc "github.com/hanherb/omdb-api/omdb-generate/contents"
)

type Movie struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

//struct for request
type MovieRequest struct {
	SearchWord string `form:"searchword" json:"searchword" binding:"required"`
	Pagination string `form:"pagination" json:"pagination" binding:"required"`
}

//struct for response
type SearchData struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

type MovieResponse struct {
	Search   []SearchData `json:"search"`
	Response string       `json:"response"`
	Error    string       `json:"error"`
}

func (m *SearchData) GrpcToMovie() *contentsGrpc.Movie {
	result := &contentsGrpc.Movie{
		Title:  m.Title,
		Year:   m.Year,
		ImdbID: m.ImdbID,
		Type:   m.Type,
		Poster: m.Poster,
	}

	return result
}
