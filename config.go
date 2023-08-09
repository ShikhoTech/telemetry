package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

// InitTracing configures an OpenTelemetry exporter and trace provider.
func InitTracing(serviceName string, exporter sdkTrace.SpanExporter) (*sdkTrace.TracerProvider, error) {
	// resource
	re := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion("v0.1.0"),
		attribute.String("library.language", "go"),
	)

	tp := sdkTrace.NewTracerProvider(
		sdkTrace.WithResource(re),
		sdkTrace.WithSampler(sdkTrace.AlwaysSample()),
		sdkTrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
