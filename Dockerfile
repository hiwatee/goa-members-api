FROM golang:1.15.3

# locale & timezone (Asia/Tokyo)
# https://github.com/moby/moby/issues/12084
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# system update
RUN apt-get update && apt-get install -y golang-goprotobuf-dev

# copy application code from host.
ADD . /go/src
WORKDIR /go/src

# install goa
RUN go get goa.design/goa/v3
RUN go get goa.design/goa/v3/...
RUN go get github.com/golang/protobuf/protoc-gen-go

# install realize
RUN GO111MODULE=off go get github.com/oxequa/realize

ENTRYPOINT ["entrypoint.sh"]