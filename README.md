# GoTool

> 一个轻量级的Go语言工具集合，提倡的核心理念为：基于Golang原生库、无第三方依赖。

## 安装

```shell
go get -u github.com/up-zero/gotool
```

## 快速开始

例如 `Md5()` 方法，使用方式如下所示，其它方法参考功能列表及其测试案例。

```go
package main

import "github.com/up-zero/gotool"
import "log"

func main() {
	data, err := gotool.Md5("123456")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
```

## 功能列表

+ **Md5** 获取md5值
+ **Md5Iterations** 迭代多次求md5
+ **Md5File** 获取文件的MD5
+ **FileCopy** 文件拷贝
+ **FileDownload** 文件下载
+ **DirCopy** 绝对目录文件拷贝
+ **FileCount** 获取指定目录下的文件个数
