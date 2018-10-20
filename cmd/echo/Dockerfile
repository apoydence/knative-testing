FROM golang

ADD . /go/src/github.com/poy/knative-testing/cmd/echo
RUN go install github.com/poy/knative-testing/cmd/echo
ENTRYPOINT /go/bin/echo

EXPOSE 8080
