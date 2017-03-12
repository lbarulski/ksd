FROM golang:1.7
RUN mkdir -p /go/src/ksd
WORKDIR /go/src/ksd
RUN curl https://glide.sh/get | sh

COPY . /go/src/ksd

RUN glide install
RUN go build

CMD ["/go/src/ksd/ksd"]