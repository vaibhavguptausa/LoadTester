package main

import (
	"fmt"
	"sync"
	"time"
)

type Strategy string

const (
	Linear      Strategy = "LINEAR"
	Exponential Strategy = "Exponential"
	Random      Strategy = "RANDOM"
)

type LoadRequest struct {
	ReqTotal int
	Time     string
	Strategy Strategy
	Request
}

func InitLoad(request LoadRequest) {
	var metricArray *[]*RequestMetric
	startTime := time.Now()
	switch request.Strategy {
	case Linear:
		metricArray, _ = SimulateLinearLoad(request)
	default:
		metricArray, _ = SimulateLinearLoad(request)
	}
	for _, val := range *metricArray {
		fmt.Printf("startTime - %v    latency- %v \n", val.startTime.Sub(startTime).Milliseconds(), val.totalDuration)
	}

}

func SimulateLinearLoad(request LoadRequest) (*[]*RequestMetric, error) {
	totalTime, err := time.ParseDuration(request.Time)
	timeBetweenReq := totalTime.Milliseconds() / int64(request.ReqTotal)
	if err != nil {
		return nil, err
	}
	metricChannel := make(chan *RequestMetric, request.ReqTotal)
	wg := sync.WaitGroup{}
	for i := 0; i < request.ReqTotal; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			metric, _ := makeRequest(Request{URL: request.URL, Method: request.Method, Body: request.Body})
			metricChannel <- metric
		}()
		time.Sleep(time.Duration(timeBetweenReq) * time.Millisecond)
	}
	wg.Wait()
	close(metricChannel)
	var metricArray []*RequestMetric
	for val := range metricChannel {
		metricArray = append(metricArray, val)
	}
	return &metricArray, nil
}
