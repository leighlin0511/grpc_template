
# build script

docker build . --target bin --output bin/ --platform linux/amd64 --build-arg BUILDKIT_INLINE_CACHE=1 -f docker/Dockerfile.build
##### go build -o bin\template-server.exe .\cmd\main.go

# run template_server

##### .\bin\template-server.exe --config=".\config\config.local.yaml" run

mac
go run cmd/main.go --config="./config/config.local.yaml" run

windows
go run cmd/main.go --config=".\config\config.local.yaml" run



# test template server using grpcurl
# Run the tool
.\grpcurl.exe --v -plaintext -d '{}' localhost:8080 service.app.OrderService/List
# test template server using curl(type "Remove-item alias:curl" if you are are using Windows)
curl http://localhost:8080/template/app/orders

.\grpcurl.exe --v -plaintext -d '{\"orderId\":\"001\"}' localhost:8080 service.app.OrderService/Retrieve
curl http://localhost:8080/template/app/order/001

.\grpcurl.exe --v -plaintext -d '{\"items\":[{\"description\":\"rubicube\",\"price\":\"32.01\"},{\"description\":\"ipad\",\"price\":\"2000.01\"}],\"paymentMethod\":{\"paymentType\":\"VISA\"}}' localhost:8080 service.app.OrderService/Create

curl -X POST http://localhost:8080/template/app/order -d '{\"items\":[{\"description\":\"rubicube\",\"price\":\"32.01\"},{\"description\":\"ipad\",\"price\":\"2000.01\"}],\"paymentMethod\":{\"paymentType\":\"VISA\"}}'

# generate protobuf golang codes
docker build . --output protobuf/generated/ --platform linux/amd64 --build-arg BUILDKIT_INLINE_CACHE=1 -f protobuf/docker/Dockerfile