/**
 ******************************************************************************
 * @file           : http.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/26
 ******************************************************************************
 */

package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}

func Post(url string, data []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}
	return body, nil
}

func GetWithQuery(baseURL string, queryParams map[string]string) ([]byte, error) {

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	query := parsedURL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	parsedURL.RawQuery = query.Encode()

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
