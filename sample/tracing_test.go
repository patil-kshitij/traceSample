package sample

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/logicmonitor/lm-telemetry-sdk-go/config"
	"github.com/logicmonitor/lm-telemetry-sdk-go/telemetry"
)

func setup() {
	ctx := context.Background()
	customAttributes := map[string]string{
		"service.namespace": "sample-namespace",
		"service.name":      "sample-service",
	}
	err := telemetry.SetupTelemetry(ctx,
		config.WithAttributes(customAttributes),
		config.WithHTTPTraceEndpoint("localhost:55681"),
	)
	if err != nil {
		log.Fatalf("error in setting up telemetry: %s", err.Error())
	}
}

func TestParentTrace(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		setup()
		ParentTrace(context.Background())
		time.Sleep(30 * time.Second)
	})
}
