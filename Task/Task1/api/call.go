package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	timelayout = "2006-01-02T15:04:05Z"
)

type (
	Response struct {
		Commit CommitDetail `json:"commit"`
		// Sha    string       `json:"sha"`
	}
	CommitDetail struct {
		Committer CommitterDetail `json:"committer"`
	}
	CommitterDetail struct {
		Date string `json:"date"`
	}
)

func Call(url, path1, path2 string) (*time.Time, error) {
	finalUrl := url + "/" + path1 + "/" + path2 + "/commits"

	resp, err := http.Get(finalUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	readBytes, _ := ioutil.ReadAll(resp.Body)

	var response []Response
	err = json.Unmarshal(readBytes, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	date := response[0].Commit.Committer.Date

	t, _ := time.Parse(timelayout, date)

	return &t, nil
}
