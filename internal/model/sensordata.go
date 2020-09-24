package model

type SensorData int
type ConvertFunc func(SensorData) SensorData

func (s SensorData) Convert(fn ConvertFunc) SensorData {
	return fn(s)
}
