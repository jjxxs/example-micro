all:
	@echo "make start-registry"
	@echo "make start-service"
	@echo "make start-client"

start-registry:
	etcd

start-service:
	MICRO_REGISTRY=etcd go run service/main.go

start-client:
	MICRO_REGISTRY=etcd go run client/main.go