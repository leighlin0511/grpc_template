echo "Generating go code"
protoc -I proto/ -I proto_vendor/ --go_out=generated/pkg --go-grpc_out==plugins=grpc:generated/pkg --grpc-gateway_out=logtostderr=true:generated/pkg ./proto/service/app/api.proto && \
protoc -I proto/ -I proto_vendor/ --go_out=generated/pkg ./proto/service/app/types.proto