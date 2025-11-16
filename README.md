<h1 align="center">
    <img width="300" src="./assets/logo.png" alt="">
</h1>

<p align="center">
   <a href="https://github.com/up-zero/gotool/fork" target="blank">
      <img src="https://img.shields.io/github/forks/up-zero/gotool?style=for-the-badge" alt="Gotool forks"/>
   </a>
   <a href="https://github.com/up-zero/gotool/stargazers" target="blank">
      <img src="https://img.shields.io/github/stars/up-zero/gotool?style=for-the-badge" alt="Gotool stars"/>
   </a>
   <a href="https://github.com/up-zero/gotool/pulls" target="blank">
      <img src="https://img.shields.io/github/issues-pr/up-zero/gotool?style=for-the-badge" alt="Gotool pull-requests"/>
   </a>
   <a href='https://github.com/up-zero/gotool/releases'>
      <img src='https://img.shields.io/github/release/up-zero/gotool?&label=Latest&style=for-the-badge'>
   </a>
</p>

一个轻量级的Go语言工具库，提倡的核心理念为：基于Golang标准库、无第三方依赖。

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
+ **FileSize** 获取文件大小

### 加解密（cryptoutil）

+ **Md5** 获取md5值
+ **Md5Iterations** 迭代多次求MD5
+ **Md5File** 获取文件的MD5
+ **HmacSHA256** 计算 SHA256
+ **HmacSHA384** 计算 HmacSHA384
+ **HmacSHA512** 计算 HmacSHA512
+ **Base64Encode** base64 编码
+ **Base64Decode** base64 解码
+ **JWTGenerate** 生成JWT
+ **JWTParse** 解析JWT
+ **Sha1** 获取SHA1值
+ **Sha256** 获取SHA256值
+ **Sha512** 获取SHA512值
+ **Sha1File** 获取文件SHA1值
+ **Sha256File** 获取文件SHA256值
+ **Sha512File** 获取文件SHA512值
+ **AesCbcEncrypt** AES CBC 加密
+ **AesCbcDecrypt** AES CBC 解密
+ **AesGcmEncrypt** AES GCM 加密
+ **AesGcmDecrypt** AES GCM 解密

### 类型转换（convertutil）

+ **StrToInt** 字符串转换为int
+ **StrToInt8** 字符串转换为int8
+ **StrToInt16** 字符串转换为int16
+ **StrToInt32** 字符串转换为int32
+ **StrToInt64** 字符串转换为int64
+ **StrToUint** 字符串转换为uint
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

+ **Abs** 绝对值
+ **Min** 最小值
+ **Max** 最大值
+ **Sum** 求和
+ **Average** 平均值
+ **ConvexHull** 计算凸包
+ **SimplifyPath** 简化路径
+ **OffsetPolygon** 多边形偏移 (内缩/外扩)
+ **TranslatePolygon** 多边形平移，按指定向量移动多边形

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
+ **FileDownload** 文件下载
+ **FileDownloadWithNotify** 带通知的文件下载
+ **FileDownloadWithProgress** 带进度的文件下载

### 系统（sysutil）

+ **ExecShell** 运行shell命令或脚本
+ **ExecCommand** 运行命令
+ **ExecShellWithNotify** 带通知的运行shell命令或脚本
+ **ExecCommandWithNotify** 带通知的运行命令
+ **ExecCommandWithOutput** 执行命令并返回合并后的标准输出和标准错误
+ **ExecCommandWithStream** 执行命令并对 stdout 和 stderr 的每一行实时调用回调函数
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
+ **NewObjectID** 生成一个新的 ObjectID

### 数组（arrayutil）

+ **Union** 数组去重，求并集
+ **Contains** 数组是否包含某个值
+ **Join** 整型拼接
+ **Concat** 数组拼接
+ **Intersect** 求多个数组的交集，数组中的元素存在时会先去重

### 图片（imageutil）

+ **Open** 打开图片
+ **Save** 保存图片
+ **Compression** 图片压缩
+ **Size** 图片尺寸
+ **GenerateCaptcha** 验证码图片生成
+ **Crop** 图片裁剪
+ **CropFile** 图片文件裁剪
+ **Resize** 图片缩放
+ **ResizeFile** 图片文件缩放
+ **Rotate** 旋转图片
+ **RotateFile** 旋转图片文件
+ **Flip** 翻转图片
+ **FlipFile** 翻转图片文件
+ **Overlay** 图片叠加
+ **OverlayFile** 图片文件叠加
+ **Grayscale** 图片灰度化
+ **GrayscaleFile** 图片文件灰度化
+ **GaussianBlur** 图片高斯模糊
+ **GaussianBlurFile** 图片文件高斯模糊
+ **AdjustBrightness** 图片亮度调整
+ **AdjustBrightnessFile** 图片文件亮度调整
+ **Invert** 图片反转颜色
+ **InvertFile** 图片文件反转颜色
+ **Binarize** 图片二值化
+ **BinarizeFile** 图片文件二值化
+ **OtsuThreshold** 基于大津法计算推荐阈值
+ **MedianBlur** 图片中值滤波
+ **MedianBlurFile** 图片文件中值滤波
+ **Sobel** 索贝尔边缘检测
+ **SobelFile** 图片文件索贝尔边缘检测
+ **NewErodeRectKernel** 创建一个用于腐蚀的矩形核
+ **NewErodeCrossKernel** 创建一个用于腐蚀的十字形核
+ **Erode** 图片腐蚀
+ **ErodeFile** 图片文件腐蚀
+ **Dilate** 图片膨胀
+ **DilateFile** 图片文件膨胀
+ **MorphologyOpen** 开运算
+ **MorphologyOpenFile** 对图片文件进行开运算
+ **MorphologyClose** 闭运算
+ **MorphologyCloseFile** 对图片文件进行闭运算
+ **EqualizeHist** 直方图均衡化
+ **EqualizeHistFile** 图片文件直方图均衡化
+ **DrawFilledCircle** 绘制填充的圆形
+ **DrawThickLine** 绘制粗线
+ **DrawLine** 绘制直线
+ **DrawRectOutline** 绘制矩形边框
+ **DrawFilledRect** 矩形填充
+ **DrawThickRectOutline** 绘制粗矩形边框
+ **DrawPolygonOutline** 绘制多边形边框
+ **DrawThickPolygonOutline** 绘制粗多边形边框
+ **DrawFilledPolygon** 多边形填充
+ **ConvexHull** 计算凸包
+ **SimplifyPath** 简化路径
+ **OffsetPolygon** 多边形偏移 (内缩/外扩)

### 条件判断（conditionutil）

+ **If** 三元运算符

### 流（streamutil）

+ **NewStream** 初始化 Stream
+ **Stream_Filter** 数据过滤
+ **Stream_Map** 数据处理
+ **Stream_Extreme** 返回流中的极值
+ **Stream_Max** 数据最大值
+ **Stream_Min** 数据最小值
+ **Stream_ToSlice** 转换为切片
+ **StreamMap** 数据处理与转换

### 结构体（structutil）

+ **SetDefaults** 设置结构体默认值
+ **NewWithDefaults** 初始化带有默认值的结构体

### 字符串（stringutil）

+ **Reverse** 字符串反转
+ **TakeFirst** 截取字符串前 n 个 Unicode 字符（rune）
+ **ContainsAny** 判断字符串 s 是否包含 substrs 中的任意一个子串
+ **ContainsAll** 判断字符串 s 是否包含 substrs 中的所有子串

### 验证器（validator）

+ **IsDigit** 判断字符串是否为数字
+ **IsAlpha** 验证字符串是否为 Unicode 字符
+ **IsAlphaStrict** 验证字符串是否为英文字符 a-z A-Z
+ **IsAlphaNumeric** 验证字符串是否为 Unicode 字符或数字
+ **IsIpv4** 验证字符串是否为 IPv4 地址
+ **IsIpv6** 验证字符串是否为 IPv6 地址

### 测试（testutil）

+ **Equal** 断言两个值相等
+ **NotEqual** 断言两个值不相等

### MAP工具（maputil）

+ **NewConcurrentMap** 初始化并发 Map
+ **ConcurrentMap_Set** 设置键值对（写操作）
+ **ConcurrentMap_Get** 获取值（读操作）
+ **ConcurrentMap_GetOrSet** 获取一个键的值（读操作），如果键不存在，则设置并返回新值（写操作）
+ **ConcurrentMap_Delete** 删除键（写操作）
+ **ConcurrentMap_GetAndDelete** 获取键的值并删除（写操作）
+ **ConcurrentMap_Len** 获取 Map 中元素的数量（读操作）
+ **ConcurrentMap_Clear** 清空 Map 中的所有元素（写操作）
+ **ConcurrentMap_Range** 遍历 Map（读操作）
