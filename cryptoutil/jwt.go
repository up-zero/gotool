package cryptoutil

import (
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/up-zero/gotool"
)

type JwtClaims struct {
	Expires int64          `json:"expires"` // 过期时间, 多少秒后过期, 0表示永不不过期
	Data    map[string]any `json:"data"`    // 数据
}

// JWTGenerate 生成JWT
//
// # Params:
//
//	claims: JWT声明
//	secretKey: 密钥
//	algorithm: 加密算法，默认为 HS256，可选值为 HS256、HS384、HS512
func JWTGenerate(claims JwtClaims, secretKey string, algorithm ...string) (string, error) {
	if claims.Expires > 0 {
		claims.Expires += time.Now().Unix()
	}

	var alg = "HS256"
	if len(algorithm) > 0 {
		alg = algorithm[0]
	}
	header := make(map[string]any)
	header["alg"] = alg
	header["typ"] = "JWT"

	// 编码header和claims
	encodedHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	encodedClaims, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// 对header和claims进行Base64编码
	base64Header := base64.RawURLEncoding.EncodeToString(encodedHeader)
	base64Claims := base64.RawURLEncoding.EncodeToString(encodedClaims)

	// 拼接header和claims，并计算签名
	unsignedToken := base64Header + "." + base64Claims
	var signature []byte
	if alg == "HS256" {
		signature = HmacSHA256([]byte(unsignedToken), []byte(secretKey))
	} else if alg == "HS384" {
		signature = HmacSHA384([]byte(unsignedToken), []byte(secretKey))
	} else if alg == "HS512" {
		signature = HmacSHA512([]byte(unsignedToken), []byte(secretKey))
	} else {
		return "", gotool.ErrInvalidJwtAlgorithm
	}

	// 拼接整个JWT
	jwt := unsignedToken + "." + base64.RawURLEncoding.EncodeToString(signature)

	return jwt, nil
}

// JWTParse 解析JWT
//
// # Params:
//
//	jwt: JWT字符串
//	secretKey: 密钥
//	algorithm: 加密算法，默认为 HS256，可选值为 HS256、HS384、HS512
func JWTParse(jwt string, secretKey string, algorithm ...string) (*JwtClaims, error) {
	// 拆分JWT为header、claims和signature
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, gotool.ErrInvalidJwtFormat
	}

	base64Header := parts[0]
	base64Claims := parts[1]
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}

	// 对claims进行Base64解码
	decodedClaims, err := base64.RawURLEncoding.DecodeString(base64Claims)
	if err != nil {
		return nil, err
	}

	// 验证签名
	var alg = "HS256"
	if len(algorithm) > 0 {
		alg = algorithm[0]
	}
	unsignedToken := base64Header + "." + base64Claims
	var expectedSignature []byte
	if alg == "HS256" {
		expectedSignature = HmacSHA256([]byte(unsignedToken), []byte(secretKey))
	} else if alg == "HS384" {
		expectedSignature = HmacSHA384([]byte(unsignedToken), []byte(secretKey))
	} else if alg == "HS512" {
		expectedSignature = HmacSHA512([]byte(unsignedToken), []byte(secretKey))
	} else {
		return nil, gotool.ErrInvalidJwtAlgorithm
	}
	if !hmac.Equal(signature, expectedSignature) {
		return nil, gotool.ErrInvalidJwtSignature
	}

	// 解析claims
	var claims = new(JwtClaims)
	err = json.Unmarshal(decodedClaims, claims)
	if err != nil {
		return nil, err
	}

	// 检查过期时间
	if claims.Expires > 0 && claims.Expires < time.Now().Unix() {
		return nil, gotool.ErrExpiredJwt
	}

	return claims, nil
}
