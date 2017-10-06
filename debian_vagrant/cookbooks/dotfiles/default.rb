git "dotfiles" do
  repository 'git://github.com/wataru0225/dotfiles.git'
end

execute 'sudo chown -R vagrant:vagrant dotfiles'
