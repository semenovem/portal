
# -----------------
# commands
# -----------------

.PHONY: help
help: Makefile
	@echo "Choose a command in:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'


## docs - собрать swagger документацию
.PHONY: docs
docs:
	@sh scripts/docs/swag.sh "swag fmt -d ./ && \
  swag init --parseDependency --parseInternal \
  -g ./cmd/api/main.go --outputTypes yaml,go --output ./docs"


## lint - go linter (wip)
.PHONY: lint
lint:
	@sh scripts/lint


## proto - compile proto
.PHONY: proto
proto:
	@sh scripts/proto/compile.sh
