from alpine:edge

RUN apk update &&\
    apk upgrade &&\
    apk add texlive &&\
    rm -rf /var/cache/apk/*
