package gotool

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// HttpGet http get 请求
//
// url: 请求地址
// header: 请求头
func HttpGet(url string, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header)
}

// HttpPost http post 请求
//
// url: 请求地址
// data: 请求参数
// header: 请求头
func HttpPost(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header)
}

// HttpPut http put 请求
//
// url: 请求地址
// data: 请求参数
// header: 请求头
func HttpPut(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header)
}

// HttpDelete http delete 请求
//
// url: 请求地址
// data: 请求参数
// header: 请求头
func HttpDelete(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", data, header)
}

// httpRequest .
func httpRequest(url, method string, data, header []byte) ([]byte, error) {
	var err error
	reader := bytes.NewBuffer(data)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// 处理 header
	if len(header) > 0 {
		headerMap := new(map[string]interface{})
		err = json.Unmarshal(header, headerMap)
		if err != nil {
			return nil, err
		}
		for k, v := range *headerMap {
			if k == "" || v == "" {
				continue
			}
			request.Header.Set(k, v.(string))
		}
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
