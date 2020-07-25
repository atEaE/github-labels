GOCMD=go

.PHONY: test run build

test:
	$(GOCMD) test ./... -v -cover

run:
	$(GOCMD) run main.go

build: 
	$(GOCMD) build -o ./dist/labels
