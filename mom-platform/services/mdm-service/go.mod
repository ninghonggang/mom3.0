module mom-platform/services/mdm-service

go 1.25

require (
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.3
	github.com/ninghonggang/mom-platform/gen v0.0.0
	github.com/ninghonggang/mom-platform/pkg v0.0.0
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.74.2
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.11
)

replace github.com/ninghonggang/mom-platform/gen => ../../gen

replace github.com/ninghonggang/mom-platform/pkg => ../../pkg

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/nats-io/nats.go v1.37.0 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/rogpeppe/go-internal v1.16.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
