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

type PostDataRequest struct {
	Data float64 `json:"data"`
}

type Response struct {
	Data any `json:"data"`
}

func append(ctx context.Context, sid string, item float64) {
	actual, _ := buffers.LoadOrStore(sid, NewSensorBuffer(10))
	sb, ok := actual.(*SensorBuffer)
	if !ok {
		return
	}

	sb.Write(item)
	fmt.Printf("wrote %f for %v\n", item, sid)
}

func getAll(ctx context.Context, sid string, resultsCh chan []float64) {
	actual, _ := buffers.LoadOrStore(sid, NewSensorBuffer(10))
	sb, ok := actual.(*SensorBuffer)
	if !ok {
		return
	}

	sb.lock.RLock()
	defer sb.lock.RUnlock()

	if sb.length == 0 {
		resultsCh <- make([]float64, 0)
		return
	}

	var results []float64 = sb.data[0:sb.length]
	resultsCh <- results
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
}

func HandleGetAllData(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")
	ctx := r.Context()
	resultCh := make(chan []float64, 1)

	go getAll(ctx, sid, resultCh)

	w.Header().Set("Content-Type", "application/json")
	var res Response
	select {
	case <-ctx.Done():
		return
	case results := <-resultCh:
		res = Response{
			Data: results,
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(res); err != nil {
				http.Error(w, "failed to encode json", http.StatusInternalServerError)
				return
			}
		}
	}
}

func ListenAndServe(port int) {
	fmt.Printf("service is listening on http://localhost:%d\n", port)
	router := http.NewServeMux()

	router.HandleFunc("POST /sensors/{id}", HandlePostData)
	router.HandleFunc("GET /sensors/{id}", HandleGetAllData)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		fmt.Printf("failed to start service: %v\n", err)
	}
}
