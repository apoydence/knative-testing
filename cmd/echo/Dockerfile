FROM golang

ADD . /go/src/github.com/apoydence/knative-testing/cmd/echo
RUN go install github.com/apoydence/knative-testing/cmd/echo
ENTRYPOINT /go/bin/echo

EXPOSE 8080
