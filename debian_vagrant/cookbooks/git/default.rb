# if you want to do that had some latest git installed,
# you should execute the command below.
# git clone git://git.kernel.org/pub/scm/git/git.git

# prepare latest git
execute 'wget https://www.kernel.org/pub/software/scm/git/git-2.10.2.tar.gz'
execute 'tar -zxf git-2.10.2.tar.gz'

# install latest git
['make configure', './configure --prefix=/usr', 'make', 'make install'].each do |cmd|
  execute cmd do 
    cwd 'git-2.10.2'
  end
end

execute 'sudo rm -f git-2.10.2.tar.gz'

# execute 'git config --global user.name wataru0225'
# execute 'git config --global user.email wataru.kikuchi0225@gmail.com'
