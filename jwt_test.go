package gotool

import "testing"

var secretKey = "get"

func TestJWTGenerate(t *testing.T) {
	claims := JwtClaims{
		Expires: 0,
		Data: map[string]any{
			"username": "get",
			"phone":    "12345678901",
		},
	}
	t.Log(JWTGenerate(claims, secretKey))
}

func TestJWTParse(t *testing.T) {
	t.Log(JWTParse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjowLCJkYXRhIjp7InBob25lIjoiMTIzNDU2Nzg5MDEiLCJ1c2VybmFtZSI6ImdldCJ9fQ.zGZ_nnZ7maC_v_ge0jpcGacQ6XE95eDMyQpKPQfJJGs", secretKey))
	t.Log(JWTParse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoxNjkyMjUwOTAwLCJkYXRhIjp7InBob25lIjoiMTIzNDU2Nzg5MDEiLCJ1c2VybmFtZSI6ImdldCJ9fQ.sqOb9Ej_cIcq2TD7w_R_tNaruRBm0IqCKKBEUGzQKLE", secretKey))
	t.Log(JWTParse("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjowLCJkYXRhIjp7InBob25lIjoiMTIzNDU2Nzg5MDEiLCJ1c2VybmFtZSI6ImdldCJ9fQ.1mWB6Wwb76fYSigDTLF_6_EeUFaYVWcyysOiF9VY_zK2wRQq_sdVEgtT6WqzZJBpGKQ3X0sZ69QeN8L568iSCA", secretKey, "HS512"))
}
