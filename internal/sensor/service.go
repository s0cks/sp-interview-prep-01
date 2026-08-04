package sensor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func EncodeJsonResponse(w http.ResponseWriter, res any) bool {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return false
	}

	return true
}

func WriteJsonResponse(w http.ResponseWriter, res any, code int) {
	w.WriteHeader(code)
	EncodeJsonResponse(w, res)
}

func WriteOkResponse(w http.ResponseWriter, data any) {
	WriteJsonResponse(w, Response{Data: data}, http.StatusOK)
}

func WriteOkStatusResponse(w http.ResponseWriter) {
	res := MessageResponse{
		Message: "ok",
	}
	WriteJsonResponse(w, res, http.StatusOK)
}

func WriteNoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func WriteAcceptedResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

func WriteNotAcceptable(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotAcceptable)
}

func CreateNewSensorBuffer() *SensorBuffer {
	total_samples := DefaultNumberOfSamples
	env_num := os.Getenv("MAX_NUMBER_OF_SAMPLES")
	if len(env_num) != 0 {
		num, err := strconv.Atoi(env_num)
		if err != nil {
			fmt.Println("conversion error:", err)
			os.Exit(1)
			return nil
		}

		total_samples = num
	}

	return NewSensorBuffer(total_samples)
}

func GetSensorBufferOrNil(sid string) (*SensorBuffer, error) {
	actual, _ := buffers.Load(sid)
	if actual == nil {
		return nil, nil
	}

	sb, ok := actual.(*SensorBuffer)
	if !ok {
		return sb, fmt.Errorf("buffer %s is nil", sid)
	}

	return sb, nil
}

func WriteData(ctx context.Context, sid string, item float64) {
	actual, _ := buffers.LoadOrStore(sid, CreateNewSensorBuffer())
	sb, ok := actual.(*SensorBuffer)
	if !ok {
		return
	}

	sb.Write(item)
	fmt.Printf("wrote %f for %v\n", item, sid)
}

type GetAllResult struct {
	Error error
	Data  []float64
}

func (r *GetAllResult) IsEmpty() bool {
	return len(r.Data) == 0
}

func GetAll(ctx context.Context, sid string, ch chan GetAllResult) {
	sb, err := GetSensorBufferOrNil(sid)
	if err != nil {
		ch <- GetAllResult{
			Error: err,
			Data:  nil,
		}
		return
	}

	if sb == nil {
		ch <- GetAllResult{
			Error: nil,
			Data:  nil,
		}
		return
	}

	var data []float64
	if sb.GetLengthSafely() == 0 {
		data = make([]float64, 0)
	} else {
		data = sb.ReadAll()
	}

	ch <- GetAllResult{
		Error: nil,
		Data:  data,
	}
}

type GetMetricResult struct {
	Error error
	Data  float64
	Empty bool
}

func GetMetric(ctx context.Context, sid string, metric string, ch chan GetMetricResult) {
	sb, err := GetSensorBufferOrNil(sid)
	if err != nil {
		ch <- GetMetricResult{
			Error: err,
			Data:  0,
		}
		return
	}

	if sb == nil {
		ch <- GetMetricResult{
			Empty: true,
		}
		return
	}

	if sb.GetLengthSafely() == 0 {
		ch <- GetMetricResult{
			Data: 0,
		}
		return
	}

	metrics, err := sb.CalcMetrics()
	if err != nil {
		ch <- GetMetricResult{
			Error: err,
		}
		return
	}

	var result float64
	switch metric {
	case "avg":
		result = metrics.avg
	case "sum":
		result = metrics.sum
	}

	ch <- GetMetricResult{
		Data: result,
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
	WriteAcceptedResponse(w)
}

func HandleGetAllData(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")
	ctx := r.Context()
	ch := make(chan GetAllResult)
	go GetAll(ctx, sid, ch)

	select {
	case <-ctx.Done():
		return
	case results := <-ch:
		if results.Error != nil {
			WriteErrorResponse(w, results.Error)
			return
		}

		if results.Data == nil || results.IsEmpty() {
			WriteNoContentResponse(w)
		} else {
			WriteOkResponse(w, results.Data)
		}
	}
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	res := ErrorResponse{
		Error: err.Error(),
	}
	w.WriteHeader(http.StatusInternalServerError)
	EncodeJsonResponse(w, res)
}

func WriteBadRequestResponse(w http.ResponseWriter, err error) {
	res := ErrorResponse{
		Error: err.Error(),
	}
	w.WriteHeader(http.StatusBadRequest)
	EncodeJsonResponse(w, res)
}

func IsValidMetric(m string) bool {
	switch m {
	case "avg", "sum":
		return true
	default:
		return false
	}
}

func HandleGetMetricData(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")
	metric := r.PathValue("metric")
	ctx := r.Context()
	ch := make(chan GetMetricResult)

	if !IsValidMetric(metric) {
		WriteBadRequestResponse(w, fmt.Errorf("'%s' metric is invalid", metric))
		return
	}

	go GetMetric(ctx, sid, metric, ch)

	select {
	case <-ctx.Done():
		return
	case val := <-ch:
		if val.Error != nil {
			WriteErrorResponse(w, val.Error)
			return
		}

		if val.Empty {
			WriteNoContentResponse(w)
			return
		}

		WriteOkResponse(w, val.Data)
	}
}

func GetAllSensorIds() []string {
	var sensors []string
	buffers.Range(func(key, v any) bool {
		if k, ok := key.(string); ok {
			sensors = append(sensors, k)
		}

		return true
	})
	return sensors
}

func HandleGetSensors(w http.ResponseWriter, r *http.Request) {
	sensors := GetAllSensorIds()
	if len(sensors) == 0 {
		WriteNoContentResponse(w)
		return
	}

	WriteOkResponse(w, sensors)
}

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	WriteOkStatusResponse(w)
}

func HandleReady(w http.ResponseWriter, r *http.Request) {
	WriteOkStatusResponse(w)
}

func ListenAndServe(port int) {
	fmt.Printf("service is listening on http://localhost:%d\n", port)
	router := http.NewServeMux()

	router.HandleFunc("POST /sensors/{id}", HandlePostData)
	router.HandleFunc("GET /sensors", HandleGetSensors)
	router.HandleFunc("GET /sensors/{id}", HandleGetAllData)
	router.HandleFunc("GET /sensors/{id}/{metric}", HandleGetMetricData)
	router.HandleFunc("GET /health", HandleHealth)
	router.HandleFunc("GET /ready", HandleReady)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		fmt.Printf("failed to start service: %v\n", err)
	}
}
