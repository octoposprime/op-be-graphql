module github.com/octoposprime/op-be-graphql

go 1.20

replace github.com/octoposprime/op-be-shared => ../op-be-shared

require (
	github.com/99designs/gqlgen v0.17.43
	github.com/golobby/container/v3 v3.3.2
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.1
	github.com/husamettinarabaci/gqltool v0.0.0-20230529191048-3a5b07a3e82e
	github.com/octoposprime/op-be-shared v0.0.0-00010101000000-000000000000
	github.com/redis/go-redis/v9 v9.4.0
	github.com/vektah/gqlparser/v2 v2.5.11
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.3 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/sosodev/duration v1.1.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
