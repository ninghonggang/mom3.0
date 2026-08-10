module mom-platform/services/trace-service

go 1.25

require (
	github.com/ninghonggang/mom-platform/gen v0.0.0
	github.com/ninghonggang/mom-platform/pkg v0.0.0
	github.com/redis/go-redis/v9 v9.5.1
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.72.2
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.11
)

replace (
	github.com/ninghonggang/mom-platform/gen => ../../gen
	github.com/ninghonggang/mom-platform/pkg => ../../pkg
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
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
	github.com/stretchr/testify v1.8.4 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
