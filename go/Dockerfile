FROM golang:1.15.3

# locale & timezone (Asia/Tokyo)
# https://github.com/moby/moby/issues/12084
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# system update
RUN apt-get update && apt-get install -y golang-goprotobuf-dev

# copy application code from host.
ADD src /go/src/members
WORKDIR /go/src/members

RUN go mod init
# RUN go build ./cmd/members

# install
# RUN go get -u goa.design/goa/v3
# RUN go get -u goa.design/goa/v3/...
# RUN go get -u github.com/golang/protobuf/protoc-gen-go

EXPOSE 8080
CMD ["./members"]