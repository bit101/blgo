default:
	@go build test.go
	@./test
	@eog test.png
