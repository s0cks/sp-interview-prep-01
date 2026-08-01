package sensor

import (
	"errors"
)

type SensorBuffer struct {
	data     []float64
	head     int
	tail     int
	length   int
	capacity int
}

func NewSensorBuffer(capacity int) *SensorBuffer {
	return &SensorBuffer{
		data:     make([]float64, capacity),
		capacity: capacity,
	}
}

func (sb *SensorBuffer) Write(val float64) {
	if sb.length == sb.capacity {
		sb.data[sb.head] = 0
		sb.head = (sb.head + 1) % sb.capacity
		sb.length--
	}

	sb.data[sb.tail] = val
	sb.tail = (sb.tail + 1) % sb.capacity
	sb.length++
}

func (sb *SensorBuffer) Read() (float64, error) {
	if sb.length == 0 {
		return 0, errors.New("buffer is empty")
	}

	next := sb.data[sb.head]
	sb.data[sb.head] = 0
	sb.head = (sb.head + 1) % sb.capacity
	sb.length--
	return next, nil
}
