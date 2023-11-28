# trjx-vehicle-api
gRPC protocol files for connecting UAV to TRJX

compile for GO


$ ls
github.com              proto_go                src                     trjx-mavlink-transfer

protoc --go_out=. --go-grpc_out=. github.com/trajectoryjp/trjx-vehicle-api/proto/*/*.proto