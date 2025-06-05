cli:
	@go build -o gonest.exe cmd/gonest/main.go cmd/gonest/generate.go cmd/gonest/new.go

run:
	@go run cmd/gonest/main.go


