package sensor

import (
	"testing"
)

func TestSensorBuffer_TestWrite(t *testing.T) {
	sb := NewSensorBuffer(5)
	v1 := 1289719712.0
	sb.Write(v1)
	if sb.length != 1 {
		t.Fatalf("expected SensorBuffer length to be 1")
	}

	if sb.head != 0 {
		t.Fatalf("expected SensorBuffer head to be 0")
	}

	if sb.tail != 1 {
		t.Fatalf("expected SensorBuffer tail to be 1")
	}
}

func TestSensorBuffer_TestRead(t *testing.T) {
	sb := NewSensorBuffer(5)
	v1 := 1289719712.0
	sb.Write(v1)
	val, err := sb.Read()
	if err != nil {
		t.Fatalf("unexpected error reading: %v", err)
	}
	if val != v1 {
		t.Errorf("expected %v, got %v", v1, val)
	}

	_, err = sb.Read()
	if err == nil {
		t.Error("expected an error reading empty buffer, received none")
	}
}

func TestSensorBuffer_TestWrappingBehavior(t *testing.T) {
	sb := NewSensorBuffer(3)
	v1 := 1289719712.0
	v2 := 128912178926.0
	v3 := 128912178926.0
	v4 := 1018983.0
	v5 := 12896158.12

	// fill buffer
	sb.Write(v1)
	sb.Write(v2)
	sb.Write(v3)

	// wrap the buffer
	sb.Write(v4)
	sb.Write(v5)

	for i, expected := range []float64{v3, v4, v5} {
		val, err := sb.Read()
		if err != nil {
			t.Fatalf("v%d failed: %v", i, err)
		}

		if val != expected {
			t.Errorf("v%d failed: expected %f got %f", i, expected, val)
		}
	}

	if sb.length != 0 {
		t.Errorf("expected an empty buffer, %d items remain in buffer", sb.length)
	}
}
