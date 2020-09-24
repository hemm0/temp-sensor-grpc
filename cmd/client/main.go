package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"tempsensor.com/v1/internal/sensorpb"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:8001", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("%v", err)
		panic("FUCKED")
	}

	defer conn.Close()

	client := sensorpb.NewSensorClient(conn)

	tempStream, err := client.TempSensor(context.Background(), &sensorpb.SensorRequest{
		ToFahrenheit: false,
	})

	if err != nil {
		panic("MORE FUCKED")
	}

	for {
		res, err := tempStream.Recv()
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		fmt.Printf("Temp: %d\n", res.Value)
	}

}
