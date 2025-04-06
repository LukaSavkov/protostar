module metrics-api

go 1.23.0

toolchain go1.23.8

require (
	github.com/c12s/magnetar v1.0.0
	github.com/gorilla/mux v1.8.1
	github.com/nats-io/nats.go v1.41.0
	github.com/robfig/cron/v3 v3.0.1
	google.golang.org/grpc v1.65.0
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/nats-io/nkeys v0.4.9 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)

replace github.com/c12s/magnetar => ../../magnetar

replace github.com/c12s/oort => ../../oort
