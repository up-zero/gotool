package netutil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
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

var httpClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,  // 连接超时
			KeepAlive: 30 * time.Second, // 保持长连接
		}).DialContext,
		MaxIdleConns:        100,              // 最大空闲连接数
		MaxIdleConnsPerHost: 100,              // 每个主机的最大空闲连接数
		IdleConnTimeout:     90 * time.Second, // 空闲连接超时时间
		TLSHandshakeTimeout: 5 * time.Second,  // TLS 握手超时
	},
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

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, method, url, reader)
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

	resp, err := httpClient.Do(request)
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

// FileDownload 文件下载
//
// # Params:
//
//	url: 文件地址
//	filePath: 文件路径
func FileDownload(url, filePath string) error {
	// 创建文件夹
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	// 创建目录源文件
	writer, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer writer.Close()
	// 下载文件
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 拷贝文件
	if _, err = io.Copy(writer, resp.Body); err != nil {
		return err
	}
	return nil
}

type DownloadProgressCallback func(finishBytes, totalBytes int64)

// DownloadProgress 包装了一个 io.Writer，并加入了进度回调功能。
type DownloadProgress struct {
	Total    int64                    // 文件总大小
	Finish   int64                    // 当前已写入大小
	callback DownloadProgressCallback // 回调函数
}

// Write 实现了 io.Writer 接口。
// 每次有数据块写入时，这个方法会被调用。
func (pw *DownloadProgress) Write(p []byte) (int, error) {
	n := len(p)
	pw.Finish += int64(n)
	// 有回调函数则调用
	if pw.callback != nil {
		pw.callback(pw.Finish, pw.Total)
	}
	return n, nil
}

// FileDownloadWithNotify 带通知的文件下载
//
// # Params:
//
//	ch: 通知进度
//	url: 文件地址
//	filePath: 文件路径
func FileDownloadWithNotify(ch chan DownloadProgress, url, filePath string) (*DownloadProgress, error) {
	defer close(ch)
	// 创建文件夹
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return nil, err
	}

	w, err := os.Create(filePath + ".tmp")
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fileSize := resp.ContentLength
	if fileSize <= 0 {
		return nil, fmt.Errorf("invalid Content-Length")
	}

	// 通知进度
	progress := &DownloadProgress{Total: fileSize}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				select {
				case ch <- *progress:
				default:
				}
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
	defer cancel()

	// 下载文件
	_, err = io.Copy(w, io.TeeReader(resp.Body, progress))
	if err != nil {
		return progress, err
	}
	w.Close()

	// 重命名文件
	if err := os.Rename(filePath+".tmp", filePath); err != nil {
		return progress, err
	}

	return progress, nil
}

// FileDownloadWithProgress 带进度的文件下载
//
// # Params:
//
//	url: 文件地址
//	filePath: 文件路径
//	callback: 用于通知下载进度
func FileDownloadWithProgress(url string, filePath string, callback DownloadProgressCallback) error {
	// 准备文件
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	w, err := os.Create(filePath + ".tmp")
	if err != nil {
		return err
	}

	// 开始下载
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	progress := &DownloadProgress{
		Total:    resp.ContentLength,
		Finish:   0,
		callback: callback,
	}
	_, err = io.Copy(w, io.TeeReader(resp.Body, progress))
	if err != nil {
		return err
	}
	w.Close()

	// 重命名文件
	if err := os.Rename(filePath+".tmp", filePath); err != nil {
		return err
	}

	return nil
}
