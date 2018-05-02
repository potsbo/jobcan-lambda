NAME := jobcanTouch

dist/$(NAME).zip: dist/$(NAME)
	cd $(dir $@) && \
	zip $(notdir $@) $(notdir $<) && \
	cd ..

dist/$(NAME): main.go vendor/*
	GOOS=linux GOARCH=amd64 go build -o $@

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep
	dep ensure -v

vendor/*: Gopkg.toml Gopkg.lock
	@make deps

.PHONY: clean
clean:
	rm -rf dist
