package netutil

import "testing"

func TestUrlBase(t *testing.T) {
	println(UrlBase("https://example.com/path/to/file.txt"))                      // file.txt
	println(UrlBase("https://example.com/path/to/file.txt?query=param#fragment")) // file.txt
}
