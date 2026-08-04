package sensor

import (
	"errors"
	"sync"
)

type SensorBuffer struct {
	lock     sync.RWMutex
	data     []float64
	head     int
	tail     int
	length   int
	capacity int
}

type SensorBufferMetrics struct {
	sum float64
	avg float64
}

const DefaultNumberOfSamples = 10

func NewSensorBuffer(capacity int) *SensorBuffer {
	return &SensorBuffer{
		data:     make([]float64, capacity),
		capacity: capacity,
	}
}

func (sb *SensorBuffer) Write(val float64) {
	sb.lock.Lock()
	defer sb.lock.Unlock()

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
	sb.lock.Lock()
	defer sb.lock.Unlock()

	if sb.length == 0 {
		return 0, errors.New("buffer is empty")
	}

	next := sb.data[sb.head]
	sb.data[sb.head] = 0
	sb.head = (sb.head + 1) % sb.capacity
	sb.length--
	return next, nil
}

func (sb *SensorBuffer) GetLengthSafely() int {
	sb.lock.RLock()
	defer sb.lock.RUnlock()
	return sb.length
}

func (sb *SensorBuffer) GetCapacitySafely() int {
	sb.lock.RLock()
	defer sb.lock.RUnlock()
	return sb.capacity
}

func (sb *SensorBuffer) CalcMetrics() (*SensorBufferMetrics, error) {
	sb.lock.RLock()
	defer sb.lock.RUnlock()

	if sb.length == 0 {
		return nil, errors.New("buffer is empty")
	}

	sum := float64(0.0)
	curr := sb.head
	for i := 0; i < sb.length; i++ {
		sum += sb.data[curr]
		curr = (curr + 1) % sb.capacity
	}

	return &SensorBufferMetrics{
		sum: sum,
		avg: sum / float64(sb.length),
	}, nil
}

func (sb *SensorBuffer) ReadAll() []float64 {
	sb.lock.RLock()
	defer sb.lock.RUnlock()

	if sb.length == 0 {
		return []float64{}
	}

	result := make([]float64, sb.length)
	curr := sb.head

	for i := 0; i < sb.length; i++ {
		result[i] = sb.data[curr]
		curr = (curr + 1) % sb.capacity
	}

	return result
}
