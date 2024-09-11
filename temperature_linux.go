//go:build linux

package gotool

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getCPUThermalZones() ([]string, error) {
	var cpuZones []string

	// 遍历 /sys/class/thermal/thermal_zone* 目录
	thermalPath := "/sys/class/thermal"
	entries, err := os.ReadDir(thermalPath)
	if err != nil {
		return nil, err
	}
	for _, v := range entries {
		path := filepath.Join(thermalPath, v.Name())
		if strings.Contains(v.Name(), "thermal_zone") {
			typePath := filepath.Join(path, "type")
			typeData, err := os.ReadFile(typePath)
			if err != nil {
				continue
			}

			// 判断传感器类型是否与 CPU 相关
			sensorType := strings.TrimSpace(string(typeData))
			if strings.Contains(strings.ToLower(sensorType), "cpu") ||
				strings.Contains(strings.ToLower(sensorType), "x86_pkg_temp") ||
				strings.Contains(strings.ToLower(sensorType), "coretemp") {
				cpuZones = append(cpuZones, path)
			}
		}
	}

	return cpuZones, nil
}

func readTemperature(tempPath string) (float64, error) {
	data, err := os.ReadFile(tempPath)
	if err != nil {
		return 0, err
	}

	temp, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, err
	}

	return float64(temp) / 1000.0, nil
}

// CPUTemperatures 获取CPU温度
func CPUTemperatures() []*TemperatureStat {
	cpuZones, err := getCPUThermalZones()
	if err != nil || len(cpuZones) == 0 {
		return nil
	}
	list := make([]*TemperatureStat, 0, len(cpuZones))

	for _, zone := range cpuZones {
		typePath := filepath.Join(zone, "type")
		tempPath := filepath.Join(zone, "temp")

		sensorType, err := os.ReadFile(typePath)
		if err != nil {
			continue
		}

		temp, err := readTemperature(tempPath)
		if err != nil {
			continue
		}

		list = append(list, &TemperatureStat{
			SensorType:  strings.TrimSpace(string(sensorType)),
			Temperature: temp,
		})
	}
	return list
}
