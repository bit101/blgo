default:
	@go build test.go
	@./test
	@open test.png
