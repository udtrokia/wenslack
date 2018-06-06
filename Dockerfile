
###
  #
  #       WENSLACK DOCKERFILE
  #
  #      udtrokia@blockheaders
  #
  #
  ###

FROM golang:1.9.6-alpine

WORKDIR /go/src/wenslack

COPY ./wenslack .

VOLUME /go/src/wenslack

EXPOSE 2333
