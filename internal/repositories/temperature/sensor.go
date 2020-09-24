package temperature

import (
	"math/rand"

	"tempsensor.com/v1/internal/model"
)

// Sensor blah
type Sensor struct {
}

// GetSensorValue blah
func (Sensor) GetSensorValue() model.SensorData {

	return model.SensorData(rand.Intn(100))
}
