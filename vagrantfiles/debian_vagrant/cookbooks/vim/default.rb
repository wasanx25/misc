execute "mkdir -p .vim/bundle"
git ".vim/bundle/neobundle.vim" do
  repository 'git://github.com/Shougo/neobundle.vim'
end
execute 'sudo chown -R vagrant:vagrant .vim'
