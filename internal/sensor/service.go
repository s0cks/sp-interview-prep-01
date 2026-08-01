package sensor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	buffers sync.Map
)

func append(ctx context.Context, sid string, item float64) {
	actual, _ := buffers.LoadOrStore(sid, NewSensorBuffer(10))
	sb, ok := actual.(*SensorBuffer)
	if !ok {
		return
	}

	sb.Write(item)
	fmt.Printf("wrote %f for %v\n", item, sid)
}

type PostDataRequest struct {
	Data float64 `json:"data"`
}

func HandlePostData(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")

	var req PostDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input request", http.StatusBadRequest)
		return
	}

	go append(context.Background(), sid, req.Data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(``))
}

func ListenAndServe(port int) {
	fmt.Printf("service is listening on http://localhost:%d\n", port)
	router := http.NewServeMux()

	router.HandleFunc("POST /sensors/{id}", HandlePostData)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		fmt.Printf("failed to start service: %v\n", err)
	}
}
