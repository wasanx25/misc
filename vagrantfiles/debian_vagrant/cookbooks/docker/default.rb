# if you want to do that had some latest docker installed,
# you should execute the command below.
# wget -N https://get.docker.com/ | sh

# install latest docker
execute 'wget -qO- https://get.docker.com/ | sh'
execute 'sudo usermod -aG docker $USER'
