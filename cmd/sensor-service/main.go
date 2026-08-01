package main

import (
	"fmt"
	"sensor-service/internal/sensor"
)

func main() {
	fmt.Println("Hello, World!")
	sensor.ListenAndServe(8080)
}
