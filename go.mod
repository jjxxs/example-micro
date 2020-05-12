module github.com/vesose/example-micro

go 1.14

require (
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/golang/protobuf v1.4.1
	github.com/google/uuid v1.1.1
	github.com/micro/go-micro/v2 v2.6.0
	github.com/micro/go-plugins/broker/nats/v2 v2.5.0
	github.com/micro/go-plugins/logger/zerolog/v2 v2.5.0
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.5.0
	github.com/micro/go-plugins/store/redis v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/store/redis/v2 v2.5.0
	github.com/rs/zerolog v1.18.0
	google.golang.org/protobuf v1.22.0
)
