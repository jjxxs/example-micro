all:
	@echo "make start-registry"
	@echo "make start-service"
	@echo "make start-client"

start-registry:
	etcd

start-counter:
	MICRO_REGISTRY=etcd go run counter-service/main.go

start-greeter:
	MICRO_REGISTRY=etcd go run greeter-service/main.go

start-client:
	MICRO_REGISTRY=etcd go run client/main.go