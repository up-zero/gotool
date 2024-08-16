package gotool

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var defaultTimeout = 5 * time.Second

// HttpGet http get 请求
//
// url: 请求地址
// header: 请求头
func HttpGet(url string, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header, defaultTimeout)
}

// HttpPost http post 请求
//
// url: 请求地址
// data: 请求参数
// header: 请求头
func HttpPost(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header, defaultTimeout)
}

// HttpPut http put 请求
//
// url: 请求地址
// data: 请求参数
// header: 请求头
func HttpPut(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header, defaultTimeout)
}

// HttpDelete http delete 请求
//
// url: 请求地址
// data: 请求参数
// header: 请求头
func HttpDelete(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", data, header, defaultTimeout)
}

// HttpGetWithTimeout http get 请求
//
// url: 请求地址
// timeout: 超时时间
// header: 请求头
func HttpGetWithTimeout(url string, timeout time.Duration, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header, timeout)
}

// HttpPostWithTimeout http post 请求
//
// url: 请求地址
// timeout: 超时时间
// data: 请求参数
// header: 请求头
func HttpPostWithTimeout(url string, data []byte, timeout time.Duration, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header, timeout)
}

// HttpPutWithTimeout http put 请求
//
// url: 请求地址
// timeout: 超时时间
// data: 请求参数
// header: 请求头
func HttpPutWithTimeout(url string, data []byte, timeout time.Duration, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header, timeout)
}

// HttpDeleteWithTimeout http delete 请求带超时时长
//
// url: 请求地址
// timeout: 超时时间
// data: 请求参数
// header: 请求头
func HttpDeleteWithTimeout(url string, data []byte, timeout time.Duration, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", data, header, timeout)
}

// httpRequest .
func httpRequest(url, method string, data, header []byte, timeout time.Duration) ([]byte, error) {
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

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second, // 连接超时为3秒
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
	request = request.WithContext(ctx)
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
