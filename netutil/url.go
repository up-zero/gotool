package netutil

import (
	"net/url"
	"path"
)

// UrlBase 获取URL路径的基础名称
//
// # Params:
//
//	rawURL: 资源的网络地址
//
// # Examples:
//
//	UrlBase("https://example.com/path/to/file.txt") // file.txt
func UrlBase(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return path.Base(u.Path)
}
