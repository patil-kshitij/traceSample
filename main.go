package main

import (
	"context"
	"log"
	"time"

	"github.com/logicmonitor/lm-telemetry-sdk-go/config"
	"github.com/logicmonitor/lm-telemetry-sdk-go/telemetry"
	"github.com/patil-kshitij/traceSample/sample"
)

var ()

func main() {
	ctx := context.Background()
	customAttributes := map[string]string{
		"service.namespace": "sample-namespace",
		"service.name":      "sample-service",
	}
	err := telemetry.SetupTelemetry(ctx,
		config.WithAttributes(customAttributes),
		config.WithHTTPTraceEndpoint("localhost:4318"),
	)
	if err != nil {
		log.Fatalf("error in setting up telemetry: %s", err.Error())
	}
	sample.ParentTrace(ctx)
	time.Sleep(30 * time.Second)
}
