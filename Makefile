gen-all:
	protoc --proto_path=proto --go_out=gen/go --go_opt=paths=import --go-grpc_out=gen/go --go-grpc_opt=paths=import proto/*.proto

gen-sso:
	protoc --proto_path=proto --go_out=gen/go --go_opt=paths=source_relative --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative proto/sso.proto

gen-orchestrator:
	protoc --proto_path=proto --go_out=gen/go --go_opt=paths=source_relative --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative proto/orchestrator.proto

gen-agent:
	protoc --proto_path=proto --go_out=gen/go --go_opt=paths=source_relative --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative proto/agent.proto
