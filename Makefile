.PHONY: bundle

build:
	go build -o gibli .

build-bin:
	go build -o /Users/mihaiblebea/go/bin/gibli .

bundle:
	@go run ./bundle/bundle.go