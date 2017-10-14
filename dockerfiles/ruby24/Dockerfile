FROM debian:stretch
ENV LANG C.UTF-8 
RUN apt-get update -y && apt-get dist-upgrade -y
RUN apt-get install -y \
      build-essential \
      libssl-dev \
      libcurl4-openssl-dev \
      libssl-dev \
      libyaml-dev \
      wget \
      zlib1g-dev

ENV RUBY_VERSION 2.4.1
ENV RUBY_FILE    ruby-$RUBY_VERSION

# Build Ruby
RUN wget http://ftp.ruby-lang.org/pub/ruby/2.4/$RUBY_FILE.tar.gz; \
    tar -zxvf $RUBY_FILE.tar.gz; \
    cd $RUBY_FILE; \
    ./configure && make && make install; \
    cd .. ; \
    rm -rf $RUBY_FILE $RUBY_FILE.tar.gz;
