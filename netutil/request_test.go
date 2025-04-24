package netutil

import (
	"bytes"
	"net/http"
	"testing"
)

type shouldBindJSONStruct struct {
	Z    string `json:"z"`
	Both string `json:"both"`
}

type shouldBindQueryBase struct {
	A string `json:"a"`
	B string `json:"b"`
}

type shouldBindQueryStruct struct {
	*shouldBindQueryBase
	Z    string  `json:"z"`
	Both string  `json:"both"`
	Num  int     `json:"num"`
	F    float64 `json:"f"`
}

// TestShouldBindJSON json入参绑定
func TestShouldBindJSON(t *testing.T) {
	r, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"z":"z","both":"both"}`)))
	if err != nil {
		t.Fatal(err)
	}
	data := new(shouldBindJSONStruct)
	err = ShouldBindJson(r, data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

// TestShouldBindQuery query入参绑定
func TestShouldBindQuery(t *testing.T) {
	r, err := http.NewRequest("GET", "/test?a=a&b=b&z=z&both=both&num=1698205401000&f=100000.121", bytes.NewBuffer([]byte(`{"z":"z","both":"both"}`)))
	if err != nil {
		t.Fatal(err)
	}
	data := new(shouldBindQueryStruct)
	err = ShouldBindQuery(r, data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
