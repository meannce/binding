start:
	helm repo add nats https://nats-io.github.io/k8s/helm/charts/
	helm repo add otel https://open-telemetry.github.io/opentelemetry-helm-charts/
	helm repo add jaeger https://jaegertracing.github.io/helm-charts
	helm repo update
	helm install nats nats/nats -f charts/nats/nats.yaml
	helm install otel otel/opentelemetry-collector -f charts/otel/otel-collector.yaml
	helm install jaeger jaeger/jaeger -f charts/jaeger/jaeger.yaml

	# optional - show default values
	helm show values nats/nats > charts/nats/default-nats.yaml
	helm show values otel/opentelemetry-collector > charts/otel/default-otel-collector.yaml

upgrade:
	helm upgrade nats nats/nats -f charts/nats/nats.yaml
	helm upgrade otel otel/opentelemetry-collector -f charts/otel/otel-collector.yaml
