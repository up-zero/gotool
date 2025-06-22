package netutil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"reflect"
	"time"
)

var defaultTimeout = 5 * time.Second

// HttpGet http get 请求
//
// # Params:
//
//	url: 请求地址
//	header: 请求头
func HttpGet(url string, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header, defaultTimeout)
}

// HttpPost http post 请求
//
// # Params:
//
//	url: 请求地址
//	data: 请求参数
//	header: 请求头
func HttpPost(url string, data any, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header, defaultTimeout)
}

// HttpPut http put 请求
//
// # Params:
//
//	url: 请求地址
//	data: 请求参数
//	header: 请求头
func HttpPut(url string, data any, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header, defaultTimeout)
}

// HttpDelete http delete 请求
//
// # Params:
//
//	url: 请求地址
//	header: 请求头
func HttpDelete(url string, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", []byte{}, header, defaultTimeout)
}

// HttpGetWithTimeout http get 请求
//
// # Params:
//
//	url: 请求地址
//	timeout: 超时时间
//	header: 请求头
func HttpGetWithTimeout(url string, timeout time.Duration, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header, timeout)
}

// HttpPostWithTimeout http post 请求
//
// # Params:
//
//	url: 请求地址
//	data: 请求参数
//	timeout: 超时时间
//	header: 请求头
func HttpPostWithTimeout(url string, data any, timeout time.Duration, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header, timeout)
}

// HttpPutWithTimeout http put 请求
//
// # Params:
//
//	url: 请求地址
//	data: 请求参数
//	timeout: 超时时间
//	header: 请求头
func HttpPutWithTimeout(url string, data any, timeout time.Duration, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header, timeout)
}

// HttpDeleteWithTimeout http delete 请求带超时时长
//
// # Params:
//
//	url: 请求地址
//	timeout: 超时时间
//	header: 请求头
func HttpDeleteWithTimeout(url string, timeout time.Duration, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", []byte{}, header, timeout)
}

// ParseResponse 解析响应
//
// # Params:
//
//	T: 为返回值类型
//	respBytes: 响应字节数组
//	err: httpRequest 错误
func ParseResponse[T any](respBytes []byte, err error) (T, error) {
	var reply T
	if err != nil {
		return reply, err
	}
	if err := json.Unmarshal(respBytes, &reply); err != nil {
		return reply, err
	}
	return reply, nil
}

// httpRequest .
func httpRequest(url, method string, data any, header []byte, timeout time.Duration) ([]byte, error) {
	var dataBytes []byte
	var err error

	// 判断 data 类型
	if reflect.TypeOf(data).Kind() == reflect.Slice && reflect.TypeOf(data).Elem().Kind() == reflect.Uint8 {
		dataBytes = data.([]byte)
	} else {
		dataBytes, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	reader := bytes.NewBuffer(dataBytes)
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

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
