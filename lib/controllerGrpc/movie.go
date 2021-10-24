package contents

import (
	"context"
	"errors"
	"strconv"

	"github.com/hanherb/omdb-api/lib/controllerRest"
	"github.com/hanherb/omdb-api/lib/models"
	contentsGrpc "github.com/hanherb/omdb-api/omdb-generate/contents"
)

func (c *Controller) MovieGetList(ctx context.Context, req *contentsGrpc.MovieGetListRequest) (*contentsGrpc.MovieGetListResponse, error) {
	//form validation
	if req.Searchword == "" {
		return nil, errors.New("searchword required")
	}
	if req.Pagination == 0 {
		return nil, errors.New("pagination required")
	}

	//convert to string
	pagination := strconv.Itoa(int(req.Pagination))

	//run search
	resp, err := controllerRest.RunSearch(ctx, &req.Searchword, &pagination)

	if err != nil {
		return nil, err
	}

	//insert log
	var qs models.MovieRequest
	qs.SearchWord = req.Searchword
	qs.Pagination = pagination

	if err := controllerRest.LogSearch(ctx, &qs, &resp.Response); err != nil {
		return nil, err
	}

	//handle gRPC response
	response := &contentsGrpc.MovieGetListResponse{}
	for _, movie := range resp.Search {
		response.Data = append(response.Data, movie.GrpcToMovie())
	}

	return response, nil
}
