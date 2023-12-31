package main

import (
	"fmt"
	"net/http"
)

func main1() {
	mock := &ClientMock{}
	err := TestMyClient(mock)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func TestMyClient(client HttpClient) error {
	request, err := http.NewRequest("GET", "https://www.reallycoolurl.com", nil)
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	fmt.Println("Successful request.")
	return nil
}

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}
