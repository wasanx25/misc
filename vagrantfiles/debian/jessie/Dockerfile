FROM debian:jessie
ENV LANG C.UTF-8 
RUN apt-get update -y && apt-get dist-upgrade -y
RUN apt-get install -y \
      sudo build-essential wget curl \
      libssl-dev \
      mysql-client \
      libcurl4-openssl-dev \
      libssl-dev \
      libmysqld-dev \
      libyaml-dev

# Set environment
ENV app     idenpo
ENV deploy  /var/local/$app
ENV user    idenpo_user
ENV RUBY_VERSION 2.3.1
ENV NODE_VERSION v6.7.0
ENV RUBY_FILE    ruby-$RUBY_VERSION
ENV NODE_FILE    node-$NODE_VERSION-linux-x64

# Build Ruby
RUN wget http://ftp.ruby-lang.org/pub/ruby/2.3/$RUBY_FILE.tar.gz; \
    tar -zxvf $RUBY_FILE.tar.gz; \
    cd $RUBY_FILE; \
    ./configure && make && make install; \
    cd .. ; \
    rm -rf $RUBY_FILE $RUBY_FILE.tar.gz;

# Build Nodejs
RUN wget -N http://nodejs.org/dist/$NODE_VERSION/$NODE_FILE.tar.gz; \
    tar xzvf $NODE_FILE.tar.gz; \
    mv $NODE_FILE /usr/local/; \
    ln -s /usr/local/$NODE_FILE/bin/npm /usr/bin/npm; \
    ln -s /usr/local/$NODE_FILE/bin/node /usr/bin/node; 

# Set User
RUN useradd -d /home/$user -m -s /bin/bash $user
RUN echo "$user:$user" | chpasswd
RUN echo "$user ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/$user
USER $user
ENV HOME /home/$user

# Working Directory
RUN sudo mkdir -p $deploy
RUN sudo chown -R $user:$user $deploy
WORKDIR $deploy

# Ruby on Rails 5
RUN sudo gem install bundler
ADD Gemfile $deploy/
ADD Gemfile.lock $deploy/
RUN sudo chown $user:$user Gemfile
RUN sudo chown $user:$user Gemfile.lock
RUN bundle install

# Angular2
ADD package.json $deploy/
ADD tsconfig.json $deploy/
ADD webpack.config.js $deploy/
RUN npm install
