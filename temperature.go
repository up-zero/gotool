package gotool

type TemperatureStat struct {
	SensorType  string  `json:"sensorType"`  // 传感器
	Temperature float64 `json:"temperature"` // 温度
}
