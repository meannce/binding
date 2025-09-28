# binding
Codename for a project where I'll have to make a small event-driven pipeline (MQTT -> NATS -> function) with OpenTelemetry spans throughout. Additionally merging some other areas I'm equally interested in exploring.

# setup
im using minikube.

`minikube start` and then `make start` does it for me. you can `make upgrade` in case you're testing changes and what not.

# todo
- probably will try to make this with helmfile at some point by curiosity

`kubectl port-forward -n default "$POD_NAME" 8080:16686` regarding Jaeger. Have to find a way to get it working without using this commmand (i guess i could add it to makefile but im stubborn)
