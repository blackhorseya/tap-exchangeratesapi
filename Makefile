.PHONY: clean
clean:
	@rm -rf coverage.txt profile.out bin ./*.csv

.PHONY: test-unit
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: lint
lint:
	@golint ./...

.PHONY: report
report:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/tap-exchangeratesapi'

.PHONY: run-to-csv
run-to-csv:
	@go run ./main.go | ~/.virtualenvs/target-csv/bin/target-csv