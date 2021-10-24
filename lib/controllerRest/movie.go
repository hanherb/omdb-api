package controllerRest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hanherb/omdb-api/lib/models"
	"github.com/hanherb/omdb-api/lib/tools"
	"github.com/hanherb/omdb-api/lib/variables"
)

func SearchMovies(c *gin.Context) {
	//get query string
	var qs models.MovieRequest

	if err := c.BindQuery(&qs); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(400, gin.H{"status": 400, "error": err.Error()})
		return
	}

	//run search
	resp, err := RunSearch(c, &qs.SearchWord, &qs.Pagination)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(400, gin.H{"status": 400, "error": err.Error()})
		return
	}

	//insert log
	if err := LogSearch(c, &qs, &resp.Response); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(400, gin.H{"status": 400, "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 200, "data": resp})
}

func RunSearch(ctx context.Context, searchWord *string, pagination *string) (*models.MovieResponse, error) {
	//creating url
	apiKey := variables.ApiKeys.OMDB
	url := "http://www.omdbapi.com/?" +
		"apikey=" + apiKey +
		"&s=" + tools.UrlEncode(searchWord) +
		"&page=" + *pagination

	//sending get request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	body, err := tools.HttpRequest(req)

	if err != nil {
		return nil, err
	}

	//handling response
	var resp models.MovieResponse
	json.Unmarshal(body, &resp)

	return &resp, nil
}

func LogSearch(ctx context.Context, qs *models.MovieRequest, status *string) error {
	searchWord := qs.SearchWord
	pagination, err := strconv.Atoi(qs.Pagination)

	if err != nil {
		return err
	}

	logger := new(models.Logger)
	logger.SearchWord = searchWord
	logger.Pagination = int32(pagination)
	logger.Status = *status

	if err := logger.Insert(ctx); err != nil {
		return err
	}

	return nil
}
