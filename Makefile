# test:
# 	go test ./...
# 
# test-coverage:
# 	go test -cover ./...
# 
# test-verbose:
# 	go test -v ./...
# 
# test-full:
# 	go test -cover -v ./...

compile:
	go build -o dist/dns-go-matic main.go
