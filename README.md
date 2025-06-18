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

import "github.com/up-zero/gotool/cryptoutil"
import "log"

func main() {
	data, err := cryptoutil.Md5("123456") // e10adc3949ba59abbe56e057f20f883e
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
```

## 方法列表

### 文件（fileutil）

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
+ **FileSync** 文件同步
+ **FileRead** 读文件（结构体）
+ **Zip** 文件夹压缩
+ **ZipWithNotify** 带通知的文件夹压缩
+ **Unzip** 文件解压
+ **UnzipWithNotify** 带通知的文件解压

### 加解密（cryptoutil）

+ **Md5** 获取md5值
+ **Md5Iterations** 迭代多次求md5
+ **Md5File** 获取文件的MD5
+ **HmacSHA256** 计算 SHA256
+ **HmacSHA384** 计算 HmacSHA384
+ **HmacSHA512** 计算 HmacSHA512
+ **JWTGenerate** 生成JWT
+ **JWTParse** 解析JWT
+ **Sha1** 获取SHA1值
+ **Sha256** 获取SHA256值
+ **Sha512** 获取SHA512值

### 类型转换（convertutil）

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
+ **CopyProperties**  复制结构体的属性

### 随机数（randomutil）

+ **RandomStr** 随机字符串
+ **RandomNumber** 随机数字
+ **RandomAlpha** 随机字母
+ **RandomAlphaNumber** 随机字母数字
+ **RandomRangeInt** 指定范围内的随机数 [最小值, 最大值)

### 数学（mathutil）

+ **Abs**  绝对值
+ **Min**  最小值
+ **Max**  最大值
+ **Sum**  求和
+ **Average**  平均值

### 网络（netutil）

+ **SendMail** 发送邮件
+ **Ipv4sLocal** 获取本地ipv4地址
+ **ShouldBindJson** json入参绑定
+ **ShouldBindQuery** query入参绑定
+ **UrlBase** 获取URL路径的基础名称
+ **HttpGet** http get 请求
+ **HttpPost** http post 请求
+ **HttpPut** http put 请求
+ **HttpDelete** http delete 请求
+ **HttpGetWithTimeout** 带超时时间的 http get 请求
+ **HttpPostWithTimeout** 带超时时间的 http post 请求
+ **HttpPutWithTimeout** 带超时时间的 http put 请求
+ **HttpDeleteWithTimeout** 带超时时间的 http delete 请求
+ **ParseResponse** 解析响应结果

### 系统（sysutil）

+ **ExecShell** 运行shell命令或脚本
+ **ExecCommand** 运行命令
+ **ExecShellWithNotify** 带通知的运行shell命令或脚本
+ **ExecCommandWithNotify** 带通知的运行命令
+ **PsByName** 根据程序名查询进程列表
+ **CPUTemperatures**  获取CPU温度
+ **SysUptime**  系统启动时间

### 日期&时间（timeutil）

+ **RFC3339ToNormalTime** RFC3339 日期格式标准化
+ **RFC1123ToNormalTime** RFC1123 日期格式标准化

### 唯一ID（idutil）

+ **UUID** uuid
+ **UUIDGenerate** 生成UUID
+ **SignalSnowflake** 单节点的雪花码

### 数组（arrayutil）

+ **Union** 数组去重，求并集
+ **Contains** 数组是否包含某个值
+ **Join** 整型拼接
+ **Concat** 数组拼接
+ **Intersect** 求多个数组的交集，数组中的元素存在时会先去重

### 图片（imageutil）

+ **ImageCompression**  图片压缩
+ **ImageSize**  图片尺寸
+ **GenerateCaptcha**  验证码图片生成

### 条件判断（conditionutil）

+ **If** 三元运算符

### 流（streamutil）

+ **NewStream** 初始化 Stream
+ **Filter** Stream方法，数据过滤
+ **Map** Stream方法，数据处理
+ **StreamMap** 数据处理与转换
+ **Extreme** Stream方法，返回流中的极值
+ **Max** Stream方法，数据最大值
+ **Min** Stream方法，数据最小值
