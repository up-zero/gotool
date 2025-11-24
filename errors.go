package gotool

import "errors"

var (
	// ErrNotSupportType 不支持的类型
	ErrNotSupportType = errors.New("not support type")

	// ErrNotIsDir 不是目录
	ErrNotIsDir = errors.New("not is dir")

	// ErrDstSrcSame 目标和源是同一个
	ErrDstSrcSame = errors.New("dst and src is same")

	// ErrInvalidJwtFormat jwt 格式错误
	ErrInvalidJwtFormat = errors.New("invalid jwt format")

	// ErrInvalidJwtSignature jwt 签名错误
	ErrInvalidJwtSignature = errors.New("invalid jwt signature")

	// ErrExpiredJwt jwt 过期
	ErrExpiredJwt = errors.New("expired jwt")
	// ErrInvalidJwtAlgorithm jwt 算法不支持
	ErrInvalidJwtAlgorithm = errors.New("invalid jwt algorithm")

	// ErrSrcDstCannotBeNil 源和目标不能为空
	ErrSrcDstCannotBeNil = errors.New("src and dst cannot be nil")

	// ErrDstMustBePointerStruct 目标必须是指针结构体
	ErrDstMustBePointerStruct = errors.New("dst must be a pointer struct")

	// ErrDstMustBePointer 目标必须为指针
	ErrDstMustBePointer = errors.New("dst must be a pointer")

	// ErrNotSupportFormat 数据格式不支持
	ErrNotSupportFormat = errors.New("not support format")

	// ErrInvalidUptimeFile uptime 文件无效
	ErrInvalidUptimeFile = errors.New("invalid uptime file")

	// ErrCannotBeEmpty 不能为空
	ErrCannotBeEmpty = errors.New("cannot be empty")

	// ErrInvalidParam 参数错误
	ErrInvalidParam = errors.New("invalid parameters")
)
