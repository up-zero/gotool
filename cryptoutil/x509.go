package cryptoutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"
)

// CertConfig 证书配置信息
type CertConfig struct {
	CommonName   string   // 证书的主域名或IP
	Organization []string // 组织名称，例如 ["UpZero"]
	Years        int      // 有效年限
	Hosts        []string // IP 或 域名
}

// GenSelfSignedCert 生成自签名证书和私钥
//
// # Params:
//
//	conf: 证书配置信息
//
// # Returns:
//
//	certPEM: 证书字符串
//	keyPEM: 私钥字符串
//	err: 错误信息
func GenSelfSignedCert(conf CertConfig) (certPEM, keyPEM string, err error) {
	if conf.Years <= 0 {
		conf.Years = 1
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}

	// 证书模板配置
	notBefore := time.Now()
	notAfter := notBefore.AddDate(conf.Years, 0, 0)
	serialNumber, _ := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: conf.Organization,
			CommonName:   conf.CommonName,
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	for _, h := range conf.Hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	// 自签名
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return "", "", err
	}

	// 编码输出
	certPEM = string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes}))
	keyPEM = string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}))

	return certPEM, keyPEM, nil
}
