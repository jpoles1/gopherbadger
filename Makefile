build:
	go build
run:
	make build && ./gopherbadger -md="README.md"
test:
	go test -v
cover:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o=coverage.html
coverage:
	make cover
configure:
	make dep
dep:
	if ! [ -x "$(command -v dep)" ]; then\
	  go get github.com/golang/dep/cmd/dep;\
	fi && dep ensure;
