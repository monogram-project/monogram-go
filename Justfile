[private]
@default:
    just --list

# Update the poetry environments, run the first time after cloning the repo.
setup:
    cd functests && poetry update

# Initialize decision records
init-decisions:
    python3 scripts/decisions.py --init

# Add a new decision record
add-decision TOPIC:
    python3 scripts/decisions.py --add "{{TOPIC}}"

jumpstart:
    sh jumpstart.sh

# Bump the version of monogram and optionally commit and push
bump BUMP:
    python3 ./scripts/bump.py --bump={{BUMP}} --save 

# Publish a new release
publish:
    python3 ./scripts/bump.py --publish

# Install both executables into your bin directory.
# You might need to adjust the destination if you have a GOBIN set.
install:
    just build-full
    just build-mini
    # Install to GOBIN if set, otherwise fallback to $(go env GOPATH)/bin
    if [ -n "$(go env GOBIN)" ]; then \
      cp monogram "$(go env GOBIN)/monogram"; \
      cp monogram-mini "$(go env GOBIN)/monogram-mini"; \
    else \
      cp monogram "$(go env GOPATH)/bin/monogram"; \
      cp monogram-mini "$(go env GOPATH)/bin/monogram-mini"; \
    fi

clean:
    # Remove the binaries
    rm -f monogram monogram-mini calculator calculator-typed monogram-test-coverage
    # Remove the test coverage reports
    rm -rf _build


# This is a recipe to build all the binaries that we can in order to
# test the build process. It is not intended to be run in production.
build-all: build-full build-mini build-for-docker build-calc

build-calc:
    # Build the calculator executable
    go build -o calculator ./cmd/calc
    # Build the typed calculator executable
    go build -o calculator-typed ./cmd/typecalc

# Alias: build both full and mini executables
build: build-full build-mini

# Build the full executable with the web server enabled but flagged for docker
build-for-docker:
    go build -tags withweb -ldflags="-X 'main.IsBuiltForDocker=true'" -o monogram ./cmd/monogram

# Build the full executable with the web server enabled
build-full:
    go build -tags withweb -o monogram ./cmd/monogram

# Build the reduced executable without the web server (and thus without the --test flag)
build-mini:
    go build -o monogram-mini ./cmd/monogram

# Run the unittests
unittest:
    # go test -v -cover ./mg
    go test -cover ./mg
    @echo

# Generate a coverage report for the unittests
unittest-coverage:
    rm -rf _build
    mkdir -p _build/
    go test -cover -coverprofile=_build/unittest.out ./mg
    go tool cover -html=_build/unittest.out -o _build/unittest.html
    #open _build/unittest.html


# Run the functional tests
functest: build-mini
    (cd functests && poetry run python3 functest.py --quiet --tests *-tests.yaml --command='../monogram-mini')

test: unittest functest

# Run the latest version of monogram and print the version
get-version:
    go run ./cmd/monogram --version

functest-coverage: build-mini
    go build -cover -o monogram-test-coverage ./cmd/monogram
    rm -rf _build
    mkdir -p _build/functest
    mkdir -p _build/merged_functest
    (cd functests && env GOCOVERDIR=../_build/functest poetry run python3 functest.py --tests *-tests.yaml --command='../monogram-test-coverage')
    go tool covdata merge -i=_build/functest -o=_build/merged_functest
    go tool covdata textfmt -i=_build/functest -o=_build/functest.out
    go tool cover -func=_build/functest.out > _build/functest.txt
    go tool cover -html=_build/functest.out -o _build/functest.html
    awk '/total:/ {print "Coverage Percentage: " $3}' _build/functest.txt
    #open _build/functest.html

