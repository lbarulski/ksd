FROM golang:1.7-onbuild
RUN curl https://glide.sh/get | sh
RUN glide install
