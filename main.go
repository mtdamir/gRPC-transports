package main

import (
	"flag"
)

func main() {

	listenAddr := flag.String("listenaddr", ":8080", "listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
