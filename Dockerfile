FROM golang:1.4

RUN go get github.com/aranair/gosnap
WORKDIR /go/src/github.com/aranair/gosnap
RUN go get ./...
RUN go install ./...

ADD configs.toml /go/bin/

WORKDIR /go/bin/
ENTRYPOINT gosnap

EXPOSE 5000
