FROM golang:1.11-stretch
LABEL maintainer=corydominguez@gmail.com

# set standard variables (this has been tested for linux only)
# for a full list of GOOS/GOARCH, see https://github.com/golang/go/blob/master/src/go/build/syslist.go
## turns the CGO off (allows for static linking so we can use the `scratch` image)
ARG CGO_ENABLED=0
## sets the CPU architecture to amd64 (windows, use: 386)
ARG GOARCH=amd64
## sets the target system to linux (windows, use: windows, mac, use: darwin)
ARG GOOS=linux

RUN apt update && \
	apt upgrade -y && \
	apt install -y \
	gcc \
	g++ \
	vim

# for developer happiness when debugging
RUN printf -- "alias ll=\"ls -al\";\nexport PATH=\"${PATH}\";" > /.profile

# set the environment variables for the build process
ENV GO111MODULE auto
ENV CGO_ENABLED ${CGO_ENABLED}
ENV GOARCH ${GOARCH}
ENV GOOS ${GOOS}

# create cache directories so this works even if consumer doesn't map volumes
RUN mkdir -p /.cache/go-build && chmod 777 -R /.cache

# famous last workds
WORKDIR /go/src/github.com/c11z/friday
CMD ["/bin/bash"]
