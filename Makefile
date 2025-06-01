cli:
	@go build -o gonest.exe cmd/gonest/main.go cmd/gonest/generate.go

run:
	@go run cmd/gonest/main.go


