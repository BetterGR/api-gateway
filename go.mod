module github.com/BetterGR/api-gateway

replace github.com/BetterGR/courses-microservice => ../courses-microservice

replace github.com/BetterGR/grades-microservice => ../grades-microservice

replace github.com/BetterGR/staff-microservice => ../staff-microservice

replace github.com/BetterGR/students-microservice => ../students-microservice

go 1.24.0

require (
	github.com/99designs/gqlgen v0.17.72
	github.com/BetterGR/courses-microservice v0.0.0-00010101000000-000000000000
	github.com/BetterGR/grades-microservice v0.0.0-20250224144127-6f9bfe793a5a
	github.com/BetterGR/staff-microservice v0.0.0-20250119121134-ebef1b46e8aa
	github.com/BetterGR/students-microservice v0.0.0-20250119121529-bf76de8e4038
	github.com/joho/godotenv v1.5.1
	github.com/vektah/gqlparser/v2 v2.5.25
	google.golang.org/grpc v1.72.0
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
