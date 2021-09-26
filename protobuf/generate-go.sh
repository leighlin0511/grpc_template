echo "Generating go code"
protoc -I proto/ --go_out=generated/pkg --go-grpc_out==plugins=grpc:generated/pkg ./proto/order.proto