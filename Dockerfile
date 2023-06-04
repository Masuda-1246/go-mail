FROM golang:1.19-alpine

RUN apk update && apk add git

# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app