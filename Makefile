
SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

TEST_FILES=$(shell find -name '*_test.go')
RESULTS_DIR=./results

TEST_OUTPUT=$(RESULTS_DIR)/test-output.xml
COV_OUTPUT=$(RESULTS_DIR)/coverage.out
COV_TXT_OUTPUT=$(RESULTS_DIR)/coverage.txt
COV_HTML_OUTPUT=$(RESULTS_DIR)/coverage.html
COB_OUTPUT=$(RESULTS_DIR)/coverage.xml
CHECK_OUTPUT=$(RESULTS_DIR)/checkstyle-result.xml

$(RESULTS_DIR):
	mkdir -p $(RESULTS_DIR)

$(COV_HTML_OUTPUT): $(COV_OUTPUT)
	go tool cover -html=$(COV_OUTPUT) -o $(COV_HTML_OUTPUT)

$(COB_OUTPUT): $(COV_OUTPUT) 
	gocover-cobertura < $(COV_OUTPUT) > $(COB_OUTPUT)


coverage: $(COV_HTML_OUTPUT) $(COB_OUTPUT) $(SOURCES)

test:
	go test -v `go list ./... | grep -v vendor` -covermode count -short


generate:
	go generate `go list ./... | grep -v vendor`


$(TEST_OUTPUT) $(COV_OUTPUT): $(RESULTS_DIR) $(SOURCES)
	go test -v -coverprofile=$(COV_OUTPUT) -covermode count -short \
		`go list ./... | grep -v vendor ` \
		2>&1 | go-junit-report > $(TEST_OUTPUT)

junit-test: $(TEST_OUTPUT)


$(CHECK_OUTPUT):  $(RESULTS_DIR) $(SOURCES)
	gometalinter -j 2 -e usr --deadline=60s --checkstyle --enable-all ./... > $(CHECK_OUTPUT) || true

vet: $(CHECK_OUTPUT)



