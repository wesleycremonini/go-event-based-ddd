package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func RemoveDuplicates[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func ArrayContains[T string | int](a []T, target T) bool {
	for _, i := range a {
		if i == target {
			return true
		}
	}
	return false
}

func CallExternalUrl(ctx context.Context, url string, headers map[string]string, method string, body io.Reader, tries int) (map[string]any, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	var response *http.Response
	for i := 0; i < tries; i++ {
		response, err = client.Do(req)
		if err == nil {
			break
		}

		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		time.Sleep(200 * time.Millisecond)
	}
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
