FROM golang

ADD . /go/src/github.com/poy/knative-testing/cmd/helloworld
RUN go install github.com/poy/knative-testing/cmd/helloworld
ENTRYPOINT /go/bin/helloworld

EXPOSE 8080
