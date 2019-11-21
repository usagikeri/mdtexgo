FROM golang:alpine as builder

COPY . /go/src/github.com/usagikeri/mdtexgo
WORKDIR /go/src/github.com/usagikeri/mdtexgo

RUN apk --update --no-cache add libpcap-dev \
                                git \
                                build-base &&\
  cd /go/src/github.com/usagikeri/mdtexgo &&\
  make deps &&\
  make build


FROM ubuntu:16.04

MAINTAINER usagikeri

WORKDIR /work

RUN apt update && apt install -y --no-install-recommends \
# latex
texlive-lang-japanese \
latexmk \
fonts-noto-cjk \
# pancoc
pandoc &&\
# cleanup
apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/github.com/usagikeri/mdtexgo/mt /usr/local/bin/mt

ENTRYPOINT ["mt"]
