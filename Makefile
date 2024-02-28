swag-init:
	swag init -g api/api.go -o api/docs
	
run: 
	go run cmd/main.go

install:
	swag init -g api/api.go -o api/docs
	go mod download
	go mod vendor
	go run cmd/main.go