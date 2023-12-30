package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
	"strconv"
	"sync"
	"time"
)

type Strategy string

const (
	Uniform     Strategy = "UNIFORM"
	Linear      Strategy = "LINEAR"
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
		metricArray, _ = SimulateUniformLoad(request)
	default:
		metricArray, _ = SimulateUniformLoad(request)
	}
	generateChart(generateChartData(*metricArray, startTime))
	fmt.Printf("done!!")
}

func generateChartData(metricArray []*RequestMetric, startTime time.Time) ([]string, []SeriesData) {
	xData := []string{}
	for _, val := range metricArray {
		xData = append(xData, fmt.Sprintf("%d", val.startTime.Sub(startTime).Milliseconds()))
	}
	latencyData := make([]opts.LineData, 0)
	firstByteData := make([]opts.LineData, 0)
	tlsDurationData := make([]opts.LineData, 0)
	connectDurationData := make([]opts.LineData, 0)
	for _, val := range metricArray {
		latencyData = append(latencyData, opts.LineData{Value: val.totalDuration})
		firstByteData = append(firstByteData, opts.LineData{Value: val.timeToFirstByte})
		tlsDurationData = append(tlsDurationData, opts.LineData{Value: val.tlsDuration})
		connectDurationData = append(connectDurationData, opts.LineData{Value: val.connectDuration})
	}
	seriesData := []SeriesData{{"latency", latencyData}, {"time to first byte", firstByteData}, {"tls duration", tlsDurationData}, {"connection duration", connectDurationData}}
	return xData, seriesData
}

func generateChart(xData []string, ySeriesData []SeriesData) {

	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Load analysis",
			Subtitle: "",
		}))

	line.SetXAxis(xData).SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	for _, val := range ySeriesData {
		line.AddSeries(val.Name, val.Data)
	}

	fileName := strconv.Itoa(int(time.Now().UnixMilli())) + "_bar.html"
	w, _ := os.Create(fileName)
	err := line.Render(w)
	if err != nil {
		return
	}

}

func SimulateUniformLoad(request LoadRequest) (*[]*RequestMetric, error) {
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

// SimulateLinearLoad  simulates the load using linear formula  Y = aX+c where x is time and y
// lets say tC is total count of the
func SimulateLinearLoad(request LoadRequest) (*[]*RequestMetric, error) {
	return nil, nil
}
