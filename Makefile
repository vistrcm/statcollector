.PHONY: lint requirements

lint:
	gometalinter --vendor ./...

requirements:
	go get -u github.com/alecthomas/gometalinter \
	&& gometalinter --install
