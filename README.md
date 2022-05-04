
# build script
go build -o bin/template_server.exe ./cmd/.

# run template_server
.\bin\template_server.exe --config="./config/config.yaml" run

windows
go run cmd/main.go --config="..\config\config.yaml" run

# test template_server
curl http://localhost:8080/orders
