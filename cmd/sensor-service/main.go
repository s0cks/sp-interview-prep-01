package main

import (
	"fmt"
	"os"
	"sensor-service/internal/sensor"
	"strconv"
)

func CheckSamplesEnvVar() int {
	env_num := os.Getenv("MAX_NUMBER_OF_SAMPLES")
	if len(env_num) == 0 {
		return sensor.DefaultNumberOfSamples
	}

	num, err := strconv.Atoi(env_num)
	if err != nil {
		fmt.Printf("value for env var MAX_NUMBER_OF_SAMPLES '%s' is not a valid int", env_num)
		os.Exit(1)
		return -1
	}

	return num
}

func main() {
	total_samples := CheckSamplesEnvVar()
	fmt.Printf("MAX_NUMBER_OF_SAMPLES=%d\n", total_samples)

	sensor.ListenAndServe(8080)
}
