//go:build linux

package sysutil

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// CPUTemperatures 获取当前 CPU 温度
func CPUTemperatures() []*TemperatureStat {
	var list []*TemperatureStat
	thermalPath := "/sys/class/thermal"

	entries, err := os.ReadDir(thermalPath)
	if err != nil {
		return nil
	}

	for _, v := range entries {
		// 仅处理 thermal_zone*
		if !strings.HasPrefix(v.Name(), "thermal_zone") {
			continue
		}

		basePath := filepath.Join(thermalPath, v.Name())

		typeData, err := os.ReadFile(filepath.Join(basePath, "type"))
		if err != nil {
			continue
		}

		sensorType := strings.TrimSpace(string(typeData))
		sensorTypeLower := strings.ToLower(sensorType)

		if isCPUSensor(sensorTypeLower) {
			tempPath := filepath.Join(basePath, "temp")

			tempData, err := os.ReadFile(tempPath)
			if err != nil {
				continue
			}

			tempRaw, err := strconv.Atoi(strings.TrimSpace(string(tempData)))
			if err != nil {
				continue
			}

			list = append(list, &TemperatureStat{
				SensorType:  sensorType,
				Temperature: float64(tempRaw) / 1000.0,
			})
		}
	}

	return list
}

func isCPUSensor(name string) bool {
	return strings.Contains(name, "cpu") ||
		strings.Contains(name, "x86_pkg_temp") ||
		strings.Contains(name, "coretemp") ||
		strings.Contains(name, "k10temp") ||
		strings.Contains(name, "acpitz")
}
