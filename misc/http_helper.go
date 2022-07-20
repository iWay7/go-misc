package misc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"os"
	"path"
	"strings"
	"time"
)

var httpClient = http.Client{
	Transport:     nil,
	CheckRedirect: nil,
	Jar:           nil,
	Timeout:       300 * time.Second,
}

func httpCallForBodyData(request *http.Request) ([]byte, error) {
	response, err := httpClient.Do(request)
	if err != nil || response == nil {
		return []byte{}, fmt.Errorf(`httpClient.Do(request): %v`, err)
	}
	defer TryClose(response.Body)
	if response.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("response.StatusCode != http.StatusOK: %v", response.StatusCode)
	}
	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("ioutil.ReadAll(response.Body): %v", err)
	}
	return bodyData, nil
}

func httpCallWithFormBodyBuildRequest(method string, url string, headerFields map[string]string, formFields map[string]string) (*http.Request, error) {
	formValues := netUrl.Values{}
	for key, value := range formFields {
		formValues.Set(key, value)
	}
	request, err := http.NewRequest(method, url, strings.NewReader(formValues.Encode()))
	if err != nil || request == nil {
		return nil, fmt.Errorf(`http.NewRequest(method, url, strings.NewReader(formValues.Encode())): %v`, err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	for key, value := range headerFields {
		request.Header.Set(key, value)
	}
	return request, nil
}

func HttpGetWithFormBodyForData(url string, headerFields map[string]string, formFields map[string]string) ([]byte, error) {
	request, err := httpCallWithFormBodyBuildRequest(http.MethodGet, url, headerFields, formFields)
	if err != nil {
		return nil, fmt.Errorf("httpCallWithFormBodyBuildRequest(http.MethodGet, url, headerFields, formFields): %v", err)
	}
	bodyData, err := httpCallForBodyData(request)
	if err != nil {
		return nil, fmt.Errorf("httpCallForBodyData(request): %v", err)
	}
	return bodyData, nil
}

func HttpGetWithFormBodyForObject(url string, headerFields map[string]string, formFields map[string]string, res interface{}) error {
	bodyData, err := HttpGetWithFormBodyForData(url, headerFields, formFields)
	if err != nil {
		return fmt.Errorf("HttpGetWithFormBodyForData(url, headerFields, formFields): %v", err)
	}
	err = json.Unmarshal(bodyData, res)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(bodyData, res): %v", err)
	}
	return nil
}

func HttpGetWithFormBodyForText(url string, headerFields map[string]string, formFields map[string]string) (string, error) {
	bodyData, err := HttpGetWithFormBodyForData(url, headerFields, formFields)
	if err != nil {
		return "", fmt.Errorf("HttpGetWithFormBodyForData(url, headerFields, formFields): %v", err)
	}
	return string(bodyData), nil
}

func HttpGetWithFormBodyForFile(url string, headerFields map[string]string, formFields map[string]string, filePath string) error {
	bodyData, err := HttpGetWithFormBodyForData(url, headerFields, formFields)
	if err != nil {
		return fmt.Errorf("HttpGetWithFormBodyForData(url, headerFields, formFields): %v", err)
	}
	directoryPath, _ := path.Split(filePath)
	err = os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm): %v", err)
	}
	err = ioutil.WriteFile(filePath, bodyData, 0666)
	if err != nil {
		return fmt.Errorf("ioutil.WriteFile(filePath, bodyData, 0666): %v", err)
	}
	return nil
}

func HttpPostWithFormBodyForData(url string, headerFields map[string]string, formFields map[string]string) ([]byte, error) {
	request, err := httpCallWithFormBodyBuildRequest(http.MethodPost, url, headerFields, formFields)
	if err != nil {
		return nil, fmt.Errorf("httpCallWithFormBodyBuildRequest(http.MethodPost, url, headerFields, formFields): %v", err)
	}
	bodyData, err := httpCallForBodyData(request)
	if err != nil {
		return nil, fmt.Errorf("httpCallForBodyData(request): %v", err)
	}
	return bodyData, nil
}

func HttpPostWithFormBodyForObject(url string, headerFields map[string]string, formFields map[string]string, res interface{}) error {
	bodyData, err := HttpPostWithFormBodyForData(url, headerFields, formFields)
	if err != nil {
		return fmt.Errorf("HttpPostWithFormBodyForData(url, headerFields, formFields): %v", err)
	}
	err = json.Unmarshal(bodyData, res)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(bodyData, res): %v", err)
	}
	return nil
}

func HttpPostWithFormBodyForText(url string, headerFields map[string]string, formFields map[string]string) (string, error) {
	bodyData, err := HttpPostWithFormBodyForData(url, headerFields, formFields)
	if err != nil {
		return "", fmt.Errorf("HttpPostWithFormBodyForData(url, headerFields, formFields): %v", err)
	}
	return string(bodyData), nil
}

func HttpPostWithFormBodyForFile(url string, headerFields map[string]string, formFields map[string]string, filePath string) error {
	bodyData, err := HttpPostWithFormBodyForData(url, headerFields, formFields)
	if err != nil {
		return fmt.Errorf("HttpPostWithFormBodyForData(url, headerFields, formFields): %v", err)
	}
	directoryPath, _ := path.Split(filePath)
	err = os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm): %v", err)
	}
	err = ioutil.WriteFile(filePath, bodyData, 0666)
	if err != nil {
		return fmt.Errorf("ioutil.WriteFile(filePath, bodyData, 0666): %v", err)
	}
	return nil
}

func httpCallWithJsonBodyBuildRequest(method string, url string, headerFields map[string]string, req interface{}) (*http.Request, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf(`json.Marshal(req): %v`, err)
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(reqData))
	if err != nil || request == nil {
		return nil, fmt.Errorf(`http.NewRequest(http.MethodPost, url, bytes.NewReader(reqData)): %v`, err)
	}
	header := request.Header
	header.Set("Content-Type", "application/json; charset=utf-8")
	for key, value := range headerFields {
		header.Set(key, value)
	}
	return request, nil
}

func HttpGetWithJsonBodyForData(url string, headerFields map[string]string, req interface{}) ([]byte, error) {
	request, err := httpCallWithJsonBodyBuildRequest(http.MethodGet, url, headerFields, req)
	if err != nil {
		return nil, fmt.Errorf("httpCallWithJsonBodyBuildRequest(http.MethodGet, url, headerFields, req): %v", err)
	}
	bodyData, err := httpCallForBodyData(request)
	if err != nil {
		return nil, fmt.Errorf("httpCallForBodyData(request): %v", err)
	}
	return bodyData, nil
}

func HttpGetWithJsonBodyForObject(url string, headerFields map[string]string, req interface{}, res interface{}) error {
	bodyData, err := HttpGetWithJsonBodyForData(url, headerFields, req)
	if err != nil {
		return fmt.Errorf("HttpGetWithJsonBodyForData(url, headerFields, req): %v", err)
	}
	err = json.Unmarshal(bodyData, res)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(bodyData, res): %v", err)
	}
	return nil
}

func HttpGetWithJsonBodyForText(url string, headerFields map[string]string, req interface{}) (string, error) {
	bodyData, err := HttpGetWithJsonBodyForData(url, headerFields, req)
	if err != nil {
		return "", fmt.Errorf("HttpGetWithJsonBodyForData(url, headerFields, req): %v", err)
	}
	return string(bodyData), nil
}

func HttpGetWithJsonBodyForFile(url string, headerFields map[string]string, formFields map[string]string, filePath string) error {
	bodyData, err := HttpGetWithJsonBodyForData(url, headerFields, formFields)
	if err != nil {
		return fmt.Errorf("HttpGetWithJsonBodyForData(url, headerFields, formFields): %v", err)
	}
	directoryPath, _ := path.Split(filePath)
	err = os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm): %v", err)
	}
	err = ioutil.WriteFile(filePath, bodyData, 0666)
	if err != nil {
		return fmt.Errorf("ioutil.WriteFile(filePath, bodyData, 0666): %v", err)
	}
	return nil
}

func HttpPostWithJsonBodyForData(url string, headerFields map[string]string, req interface{}) ([]byte, error) {
	request, err := httpCallWithJsonBodyBuildRequest(http.MethodPost, url, headerFields, req)
	if err != nil {
		return nil, fmt.Errorf("httpCallWithJsonBodyBuildRequest(http.MethodPost, url, headerFields, req): %v", err)
	}
	bodyData, err := httpCallForBodyData(request)
	if err != nil {
		return nil, fmt.Errorf("httpCallForBodyData(request): %v", err)
	}
	return bodyData, nil
}

func HttpPostWithJsonBodyForObject(url string, headerFields map[string]string, req interface{}, res interface{}) error {
	bodyData, err := HttpPostWithJsonBodyForData(url, headerFields, req)
	if err != nil {
		return fmt.Errorf("HttpPostWithJsonBodyForData(url, headerFields, req): %v", err)
	}
	err = json.Unmarshal(bodyData, res)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(bodyData, res): %v", err)
	}
	return nil
}

func HttpPostWithJsonBodyForText(url string, headerFields map[string]string, req interface{}) (string, error) {
	bodyData, err := HttpPostWithJsonBodyForData(url, headerFields, req)
	if err != nil {
		return "", fmt.Errorf("HttpPostWithJsonBodyForData(url, headerFields, req): %v", err)
	}
	return string(bodyData), nil
}

func HttpPostWithJsonBodyForFile(url string, headerFields map[string]string, formFields map[string]string, filePath string) error {
	bodyData, err := HttpPostWithJsonBodyForData(url, headerFields, formFields)
	if err != nil {
		return fmt.Errorf("HttpPostWithJsonBodyForData(url, headerFields, formFields): %v", err)
	}
	directoryPath, _ := path.Split(filePath)
	err = os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm): %v", err)
	}
	err = ioutil.WriteFile(filePath, bodyData, 0666)
	if err != nil {
		return fmt.Errorf("ioutil.WriteFile(filePath, bodyData, 0666): %v", err)
	}
	return nil
}
