echo "Generating go code"
protoc -I proto/ -I vendor_proto/ --go_out=generated/pkg --go-grpc_out==plugins=grpc:generated/pkg --grpc-gateway_out=logtostderr=true:generated/pkg ./proto/order.proto