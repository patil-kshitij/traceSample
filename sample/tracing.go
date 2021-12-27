package sample

import (
	"fmt"

	"github.com/patil-kshitij/traceSample/greeter"
	"go.opentelemetry.io/otel"
	"golang.org/x/net/context"
)

func ParentTrace(ctx context.Context) {
	tracer := otel.Tracer("parent-tracer")
	ctx, span := tracer.Start(ctx, "ParentTrace")
	defer span.End()
	fmt.Println("Message::", greeter.Greet())
	ChildTrace(ctx)
}

func ChildTrace(ctx context.Context) {
	tracer := otel.Tracer("child-tracer")
	_, span := tracer.Start(ctx, "ChildTrace")
	defer span.End()
}
