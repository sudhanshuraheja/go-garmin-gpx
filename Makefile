test:
	go test ./... -timeout 15s -cover -coverprofile=coverage.out -v
	go tool cover -html=coverage.out -o coverage.html

upgrade:
	go get -u all
	go mod tidy