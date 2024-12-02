
all: 
	@mkdir -p build
	@go build -o build/foo-cni cmd/main.go
