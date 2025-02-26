# GoTool

> 一个轻量级的Go语言工具库，提倡的核心理念为：基于Golang原生库、无第三方依赖。

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

## 方法列表

+ **Md5** 获取md5值
+ **Md5Iterations** 迭代多次求md5
+ **Md5File** 获取文件的MD5
+ **FileCopy** 文件拷贝
+ **FileMove** 文件移动
+ **FileDownload** 文件下载
+ **FileDownloadWithNotify** 带通知的文件下载
+ **DirCopy** 绝对目录文件拷贝
+ **CurrentDirCount** 当前文件夹下(不迭代子文件夹)文件或文件夹的个数
+ **MkParentDir** 创建父级文件夹
+ **FileCount** 获取指定目录下的文件个数
+ **FileMainName** 获取指定路径的文件名
+ **FileSave** 保存文件
+ **Zip** 文件夹压缩
+ **ZipWithNotify** 带通知的文件夹压缩
+ **Unzip** 文件解压
+ **UnzipWithNotify** 带通知的文件解压
+ **ExecShell** 运行shell命令或脚本
+ **ExecCommand** 运行命令
+ **ExecShellWithNotify** 带通知的运行shell命令或脚本
+ **ExecCommandWithNotify** 带通知的运行命令
+ **PsByName** 根据程序名查询进程列表
+ **HmacSHA256** 计算 SHA256
+ **HmacSHA384** 计算 HmacSHA384
+ **HmacSHA512** 计算 HmacSHA512
+ **JWTGenerate** 生成JWT
+ **JWTParse** 解析JWT
+ **RFC3339ToNormalTime** RFC3339 日期格式标准化
+ **RFC1123ToNormalTime** RFC1123 日期格式标准化
+ **If** 三元运算符
+ **ArrayDuplicate** 数组去重
+ **ArrayIn** 数组是否包含某个值
+ **ArrayJoin** 整型拼接
+ **RandomStr** 随机字符串
+ **RandomNumber** 随机数字
+ **RandomAlpha** 随机字母
+ **RandomAlphaNumber** 随机字母数字
+ **RandomRangeInt** 指定范围内的随机数 [最小值, 最大值)
+ **StrToInt8** 字符串转换为int8
+ **StrToInt16** 字符串转换为int16
+ **StrToInt32** 字符串转换为int32
+ **StrToInt64** 字符串转换为int64
+ **StrToUint8** 字符串转换为uint8
+ **StrToUint16** 字符串转换为uint16
+ **StrToUint32** 字符串转换为uint32
+ **StrToUint64** 字符串转换为uint64
+ **StrToFloat32** 字符串转换为float32
+ **StrToFloat64** 字符串转换为float64
+ **Int64ToStr** int64转换为字符串
+ **Uint64ToStr** uint64转换为字符串
+ **Float64ToStr** float64转换为字符串
+ **Int64ToHex** int64转换为十六进制字符串
+ **HexToInt64** 十六进制字符串转换为int64
+ **SendMail** 发送邮件
+ **Ipv4sLocal** 获取本地ipv4地址
+ **ShouldBindJson** json入参绑定
+ **ShouldBindQuery** query入参绑定
+ **UUID** uuid
+ **UUIDGenerate** 生成UUID
+ **SignalSnowflake** 单节点的雪花码
+ **UrlBase**  获取URL路径的基础名称
+ **CPUTemperatures**  获取CPU温度
+ **CopyProperties**  复制结构体的属性
+ **ImageCompression**  图片压缩
+ **ImageSize**  图片尺寸
+ **GenerateCaptcha**  验证码图片生成
+ **MathAbs**  绝对值
+ **MathMin**  最小值
+ **MathMax**  最大值
