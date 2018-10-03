FROM golang

ADD . /go/src/github.com/apoydence/knative-testing/cmd/helloworld
RUN go install github.com/apoydence/knative-testing/cmd/helloworld
ENTRYPOINT /go/bin/helloworld

EXPOSE 8080
