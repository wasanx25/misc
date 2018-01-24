FROM debian:stretch

RUN apt-get update && apt-get install -y \
      g++ gcc git libc6-dev make pkg-config tar wget

ENV app    golang_app
ENV deploy /var/www/$app

ENV GO_VERSION 1.9
ENV OS_TYPE    linux
ENV OS_ARCH    amd64

RUN wget https://storage.googleapis.com/golang/go$GO_VERSION.$OS_TYPE-$OS_ARCH.tar.gz; \
      tar -C /usr/local -xzf go$GO_VERSION.$OS_TYPE-$OS_ARCH.tar.gz

ENV PATH   $PATH:/usr/local/go/bin

RUN mkdir -p $deploy
WORKDIR $deploy
