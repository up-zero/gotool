package fileutil

import (
	"testing"
)

func TestZip(t *testing.T) {
	t.Log(Zip("E:\\github\\up-zero\\gotool\\tmp\\a", "E:\\github\\up-zero\\gotool\\tmp\\dir.zip"))
}

func TestZipWithNotify(t *testing.T) {
	dest := "E:\\github\\up-zero\\gotool\\tmp\\dir.zip"
	src := "E:\\github\\up-zero\\gotool\\tmp\\a"
	fileCnt, _ := FileCount(src)
	ch := make(chan int)
	defer close(ch)
	go func() {
		for {
			select {
			case index := <-ch:
				t.Log("progress => ", index)
				if index+1 >= fileCnt {
					return
				}
			}
		}
	}()
	t.Log(ZipWithNotify(src, dest, ch))
}

func TestUnzip(t *testing.T) {
	t.Log(Unzip("E:\\github\\up-zero\\gotool\\tmp\\dir.zip", "E:\\github\\up-zero\\gotool\\tmp\\b"))
}

func TestUnzipWithNotify(t *testing.T) {
	src := "E:\\github\\up-zero\\gotool\\tmp\\dir.zip"
	dest := "E:\\github\\up-zero\\gotool\\tmp\\b"
	ch := make(chan *UnzipNotify)
	defer close(ch)
	finish := make(chan struct{})
	defer close(finish)
	go func() {
		for {
			select {
			case data := <-ch:
				t.Log("data => ", data)
			case <-finish:
				return
			}
		}
	}()
	if err := UnzipWithNotify(src, dest, ch); err != nil {
		t.Fatal(err)
	}
	finish <- struct{}{}
}
