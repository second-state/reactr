module github.com/suborbital/reactr

go 1.17

require (
	github.com/bytecodealliance/wasmtime-go v0.31.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/pkg/errors v0.9.1
	github.com/second-state/WasmEdge-go v0.9.0-rc5
	github.com/suborbital/grav v0.4.1
	github.com/suborbital/vektor v0.5.1-0.20211112160641-0b7e68b46795
	github.com/wasmerio/wasmer-go v1.0.4
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.9.1 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/sethvargo/go-envconfig v0.4.0 // indirect
	golang.org/x/net v0.0.0-20211201190559-0a0e4e1bb54c // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/bytecodealliance/wasmtime-go => github.com/suborbital/wasmtime-go v0.31.0-subo
