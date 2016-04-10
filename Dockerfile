FROM golang:1.4

RUN get bitbucket.org/liamstask/goose/cmd/goose
RUN go get github.com/aranair/gosnap

WORKDIR /go/src/github.com/aranair/gosnap
RUN go get ./...
RUN go install ./...

ADD configs.toml /go/bin/

ENTRYPOINT /go/bin/gosnap

EXPOSE 5000
