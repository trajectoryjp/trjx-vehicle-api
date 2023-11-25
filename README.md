# trjx-vehicle-api
gRPC protocol files for connecting UAV to TRJX

compile for GO

$ ls
LICENSE         README.md       proto           proto_go

$protoc --go_out=. --go-grpc_out=. proto/*/*.proto