.PHONY: clean testing

test:
	go test ./service/... -v -cover