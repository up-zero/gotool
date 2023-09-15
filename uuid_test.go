package gotool

import (
	"log"
	"testing"
)

func TestUUID(t *testing.T) {
	l := 100000
	g := 1000
	arr := make([]string, 0, l*g)
	c := make(chan struct{})
	for j := 0; j < g; j++ {
		go func() {
			for i := 0; i < l; i++ {
				str, _ := UUID()
				arr = append(arr, str)
				c <- struct{}{}
			}
		}()
	}
	t1 := 0
	for {
		<-c
		t1++
		if t1 == g {
			break
		}
	}

	m := make(map[string]struct{})
	for _, v := range arr {
		if _, ok := m[v]; ok {
			t.Fatal(v)
		}
		m[v] = struct{}{}
	}
	log.Println("pass")
}
