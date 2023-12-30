package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httptrace"
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

type Request struct {
	URL    string
	Method string
	Body   []byte
}

func makeRequest(request Request) (*RequestMetric, error) {
	var req *http.Request
	startTime := time.Now()
	switch request.Method {
	case http.MethodGet:
		req, _ = http.NewRequest(request.Method, request.URL, nil)
	case http.MethodPost:
		reader := bytes.NewReader(request.Body)
		req, _ = http.NewRequest(request.Method, request.URL, reader)
	default:
		//return nil, errors.New("method not supported yet")
	}

	var start, connect, dns, tlsHandshake time.Time
	var dnsDuration, tlsDuration, connectDuration, timeToFirstByte, totalDuration int64

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			dnsDuration = time.Since(dns).Milliseconds()
		},

		TLSHandshakeStart: func() { tlsHandshake = time.Now() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			tlsDuration = time.Since(tlsHandshake).Milliseconds()
		},

		ConnectStart: func(network, addr string) { connect = time.Now() },
		ConnectDone: func(network, addr string, err error) {
			connectDuration = time.Since(connect).Milliseconds()
		},

		GotFirstResponseByte: func() {
			timeToFirstByte = time.Since(start).Milliseconds()
		},
	}
	ctx := httptrace.WithClientTrace(req.Context(), trace)
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	start = time.Now()
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Fatal(err)
	}
	totalDuration = time.Since(start).Milliseconds()
	metric := &RequestMetric{
		totalDuration:   totalDuration,
		timeToFirstByte: timeToFirstByte,
		connectDuration: connectDuration,
		tlsDuration:     tlsDuration,
		dnsDuration:     dnsDuration,
		startTime:       startTime,
	}
	return metric, nil
}
