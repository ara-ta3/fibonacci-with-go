run:
	go run ./fib.go

test:
	go test -v ./...

bench:
	go test -bench .
