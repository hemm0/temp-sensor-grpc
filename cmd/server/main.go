package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"tempsensor.com/v1/internal/model"
	"tempsensor.com/v1/internal/repositories/temperature"

	"tempsensor.com/v1/internal/sensorpb"

	"google.golang.org/grpc"
)

type server struct {
	humidity    sensor
	temperature sensor
}

type sensor interface {
	GetSensorValue() model.SensorData
}

func CelciusToFahrenheit(s model.SensorData) model.SensorData {
	//(0°C × 9/5) + 32
	return (s * 9 / 5) + 32
}

func (s *server) TempSensor(r *sensorpb.SensorRequest, stream sensorpb.Sensor_TempSensorServer) error {
	//talk to model to give me data

	for {
		time.Sleep(3 * time.Second)
		val := s.temperature.GetSensorValue()

		if r.ToFahrenheit {
			val = val.Convert(CelciusToFahrenheit)
		}

		sr := sensorpb.SensorResponse{Value: int64(val)}
		if err := stream.Send(&sr); err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}

	}

}

func (*server) HumiditySensor(*sensorpb.SensorRequest, sensorpb.Sensor_HumiditySensorServer) error {

	return nil
}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:8001")

	if err != nil {
		log.Fatalf("Error establising Listener : %v", err)
	}

	srv := grpc.NewServer()

	s := &server{
		temperature: temperature.Sensor{},
	}

	sensorpb.RegisterSensorServer(srv, s)

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Error establishing Server : %v", err)
	}
}
