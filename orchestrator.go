package main

import (
	"fmt"
	"sync"
	"time"
)

type Strategy string

const (
	Uniform     Strategy = "UNIFORM"
	Exponential Strategy = "Exponential"
	Random      Strategy = "RANDOM"
)

func InitLoad(request LoadRequest) {
	var metricArray *[]*RequestMetric
	startTime := time.Now()
	if request.Strategy == "" {
		request.Strategy = Uniform
	}

	switch request.Strategy {
	case Uniform:
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
