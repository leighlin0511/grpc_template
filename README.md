
# build script

docker build . --target bin --output bin/ --platform linux/amd64 --build-arg BUILDKIT_INLINE_CACHE=1 -f docker/Dockerfile.build
##### go build -o bin/template-server.exe ./cmd/.

# run template_server

##### .\bin\template_server.exe --config="./config/config.yaml" run

mac
go run cmd/main.go --config="./config/config.yaml" run

windows
go run cmd/main.go --config=".\config\config.yaml" run

# test template server
curl http://localhost:8080/orders
