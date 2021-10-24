package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(r *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(r)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return body, nil
}
