package gotool

import "testing"

func TestZip(t *testing.T) {
	t.Log(Zip("E:\\github\\up-zero\\gotool\\tmp\\dir.zip", "E:\\github\\up-zero\\gotool\\tmp\\a"))
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
	t.Log(ZipWithNotify(dest, src, ch))
}
