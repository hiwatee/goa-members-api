FROM golang:1.15.3

# locale & timezone (Asia/Tokyo)
# https://github.com/moby/moby/issues/12084
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# system update
RUN apt-get update && \
    apt-get install -y vim less

# copy application code from host.
ADD src /go/src/api
WORKDIR /go/src/api

RUN go mod init
RUN go get -u goa.design/goa/v3
RUN go get -u goa.design/goa/v3/...

# install packages. (vendoring: /go/bin/api/vendor/)
#RUN dep ensure -v -vendor-only=true

# install development tools.
#RUN go install api/vendor/github.com/goadesign/goa/goagen

# build
#RUN go build -o vendor/server

#EXPOSE 8080
#CMD ["./vendor/server"]