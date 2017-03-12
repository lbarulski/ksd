FROM golang:1.7-onbuild
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep status
RUN dep ensure