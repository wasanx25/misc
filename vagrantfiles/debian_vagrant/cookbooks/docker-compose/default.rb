# install latest docker-compose
execute 'sudo -i'
execute 'curl -L "https://github.com/docker/compose/releases/download/1.8.1/docker-compose-$(uname -s)-$(uname -m)" > /usr/local/bin/docker-compose'
execute 'chmod +x /usr/local/bin/docker-compose'
execute 'exit'
