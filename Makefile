build:
	go build
run:
	make build && ./gopherbadger
test:
	go test -v
cover:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o=coverage.html
coverage:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o=coverage.html
configure:
	make dep
dep:
	if ! [ -x "$(command -v dep)" ]; then\
	  go get github.com/golang/dep/cmd/dep;\
	fi && dep ensure;
