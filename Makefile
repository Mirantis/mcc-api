.DEFAULT_GOAL=all

LOCAL_BIN:=$(CURDIR)/bin
BIN_LINTER:=${LOCAL_BIN}/golangci-lint
EXAMPLE_DIR:=$(CURDIR)/example

.PHONY:all
all: lint go-example-test go-example-build

.PHONY: modules
modules: ## Runs go mod to ensure modules are up to date.
	# API project
	go mod tidy
	go mod vendor

	# Example project
	(cd ${EXAMPLE_DIR}; go mod tidy; go mod vendor)

.PHONY:go-test
go-example-test:
	(cd ${EXAMPLE_DIR}; go test -count=1 -v ./...)

.PHONY:go-build
go-example-build: EXAMPLE_BIN=${LOCAL_BIN}/example
go-example-build:
	rm -f ${EXAMPLE_BIN}
	(cd ${EXAMPLE_DIR}; go build -o ${EXAMPLE_BIN} main.go)

.PHONY:lint
lint: install-lint
	${BIN_LINTER} run -v --color=always ./...
	(cd ${EXAMPLE_DIR}; ${BIN_LINTER} run -v --color=always ./...)

.PHONY: install-lint
install-lint: LINTER_VERSION:=v1.43.0
install-lint:
	$(call fn_install_gotool,github.com/golangci/golangci-lint/cmd/golangci-lint,${LINTER_VERSION},${BIN_LINTER})

# fn_install_gotool installs tool from remote repository
# params:
# 1. remote repository URL
# 2. tag/branch
# 3. path to binary file
# 4. build properties
define fn_install_gotool
	@[ ! -f ${3}@${2} ] \
		|| exit 0 \
		&& echo "Installing ${1} ..." \
		&& tmp=$(shell mktemp -d) \
		&& cd "$$tmp" \
		&& echo "Tool: ${1}" \
		&& echo "Version: ${2}" \
		&& echo "Binary: ${3}" \
		&& echo "Temp: $$tmp" \
		&& go mod init temp && go get -d ${1}@${2} && go build ${4} -o ${3}@${2} ${1} \
		&& ln -sf ${3}@${2} ${3} \
		&& rm -rf "$$tmp" \
		&& echo "success istalled: ${3}" \
		&& echo "********************************************************"
endef
