install: test
	go install .

test:
	go test ./...