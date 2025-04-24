package sysutil

import (
	"fmt"
	"testing"
)

func TestCPUTemperatures(t *testing.T) {
	temps := CPUTemperatures()
	for _, v := range temps {
		fmt.Printf("CPU 传感器 (%s) 温度: %.1f°C\n", v.SensorType, v.Temperature)
	}
}
