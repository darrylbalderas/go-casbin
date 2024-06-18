lint:
	@go fmt ./...
install:
	@go mod tidy
	@go get github.com/casbin/casbin/v2
gocasbin:
	@go build -o gocasbin .
	@./gocasbin