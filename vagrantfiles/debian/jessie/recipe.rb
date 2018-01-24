def include_cookbook(name)
  root_dir = File.expand_path('../', __FILE__)
  include_recipe File.join(root_dir, 'cookbooks', name, 'default')
end

include_cookbook 'git'
include_cookbook 'docker'
# include_cookbook 'docker-compose'
include_cookbook 'dotfiles'
include_cookbook 'vim'
include_cookbook 'zsh'
include_cookbook 'working_directory'
