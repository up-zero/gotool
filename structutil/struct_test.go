package structutil

import "testing"

func TestToMap(t *testing.T) {
	type Base struct {
		ID int `json:"id"`
	}
	type Info struct {
		Age int `json:"age"`
	}
	type User struct {
		Base
		Name string `json:"name"`
		Info Info   `json:"info"`
	}
	u := User{Base: Base{ID: 1}, Name: "Tom", Info: Info{Age: 18}}
	m, _ := ToMap(u)
	t.Logf("m: %+v", m)
	//m: {"id": 1, "name": "Tom", "info": {"age": 18}}
}
