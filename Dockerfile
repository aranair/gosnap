FROM golang:onbuild

RUN go get bitbucket.org/liamstask/goose/cmd/goose

RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]

ADD configs.toml /go/bin/
ADD dbconf.yml /go/src/app/db/

EXPOSE 5000
