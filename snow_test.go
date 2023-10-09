package gotool

import (
	"fmt"
	"testing"
)

func TestSignalSnowflake(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := SignalSnowflake()
		fmt.Println(id)
	}
}
