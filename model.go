package main

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"time"
)

type RequestMetric struct {
	connectDuration int64
	timeToFirstByte int64
	tlsDuration     int64
	totalDuration   int64
	dnsDuration     int64
	statusCode      int
	startTime       time.Time
}

type LoadRequest struct {
	ReqTotal int      `json:"req_total" validate:"required"`
	Time     string   `json:"time" validate:"required"`
	Strategy Strategy `json:"strategy"`
	Request
}

type Request struct {
	URL    string `json:"url" validate:"required"`
	Method string `json:"method" validate:"required"`
	Body   string `json:"body"`
}

type JsonResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SeriesData struct {
	Name string
	Data []opts.LineData
}
