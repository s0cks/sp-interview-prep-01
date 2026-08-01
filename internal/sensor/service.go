package sensor

import (
	"fmt"
	"net/http"
	"sync"
)

type SensorService struct {
	buffers sync.Map
}

func (s *SensorService) ListenAndServe(port int) {
	router := http.NewServeMux()
	fmt.Printf("service is listening on http://localhost:%d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		fmt.Printf("failed to start service: %v\n", err)
	}
}
