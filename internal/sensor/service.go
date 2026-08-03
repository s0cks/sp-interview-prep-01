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

func WriteData(ctx context.Context, sid string, item float64) {
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

func getMetric(ctx context.Context, sid string, metric string, resultsCh chan any) {
	actual, _ := buffers.LoadOrStore(sid, NewSensorBuffer(10))
	sb, ok := actual.(*SensorBuffer)
	if !ok {
		resultsCh <- fmt.Errorf("invalid buffer type")
		return
	}

	sb.lock.RLock()
	defer sb.lock.RUnlock()
	if sb.length == 0 {
		resultsCh <- nil
		return
	}

	results, err := sb.CalcMetrics()
	if err != nil {
		resultsCh <- err
		return
	}

	switch metric {
	case "avg":
		resultsCh <- results.avg
	case "sum":
		resultsCh <- results.sum
	}
}

func HandlePostData(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")

	var req PostDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input request", http.StatusBadRequest)
		return
	}

	go WriteData(context.Background(), sid, req.Data)

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

type ErrorResponse struct {
	Error error
}

type MessageResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func HandleGetMetricData(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")
	metric := r.PathValue("metric")
	ctx := r.Context()
	resultCh := make(chan any)

	fmt.Printf("getting %s metric for sensor %s\n", metric, sid)

	w.Header().Set("Content-Type", "application/json")
	switch metric {
	case "avg", "sum":
		go getMetric(ctx, sid, metric, resultCh)
	default:
		w.WriteHeader(http.StatusBadRequest)
		var payload MessageResponse = MessageResponse{
			Message: fmt.Sprintf("invalid metric: %s", metric),
		}

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, "failed to encode json", http.StatusInternalServerError)
			return
		}
		return
	}

	var res Response
	select {
	case <-ctx.Done():
		return
	case val := <-resultCh:
		switch v := val.(type) {
		case nil:
			w.WriteHeader(http.StatusNoContent)
			return
		case error:
			w.WriteHeader(http.StatusInternalServerError)
			payload := struct {
				Error string `json:"message"`
			}{
				Error: v.Error(),
			}

			if err := json.NewEncoder(w).Encode(payload); err != nil {
				http.Error(w, "failed to encode json", http.StatusInternalServerError)
				return
			}
		default:
			res = Response{
				Data: v,
			}

			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(res); err != nil {
				http.Error(w, "failed to encode json", http.StatusInternalServerError)
				return
			}
		}
	}
}

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := StatusResponse{
		Status: "OK",
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func HandleReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := StatusResponse{
		Status: "OK",
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func ListenAndServe(port int) {
	fmt.Printf("service is listening on http://localhost:%d\n", port)
	router := http.NewServeMux()

	router.HandleFunc("POST /sensors/{id}", HandlePostData)
	router.HandleFunc("GET /sensors/{id}", HandleGetAllData)
	router.HandleFunc("GET /sensors/{id}/{metric}", HandleGetMetricData)
	router.HandleFunc("GET /health", HandleHealth)
	router.HandleFunc("GET /ready", HandleReady)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		fmt.Printf("failed to start service: %v\n", err)
	}
}
