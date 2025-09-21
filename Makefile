start:
	helm install nats ./charts/nats
	helm install otel ./charts/otel-collector
