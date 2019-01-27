IMAGE_TAG=friday:latest
UGID=$(shell id -u):$(shell id -g)

build_quiet:
	@docker build \
		--quiet \
		--tag=$(IMAGE_TAG) \
		.

build:
	@docker build \
		--tag=$(IMAGE_TAG) \
		.

console: build_quiet
	@docker run \
		--rm \
		--tty \
		--interactive \
		--user $(UGID) \
		--volume $(CURDIR):/src/github.com/c11z/friday \
		--env-file $(CURDIR)/env.list \
		--workdir /src/github.com/c11z/friday \
		$(IMAGE_TAG)
