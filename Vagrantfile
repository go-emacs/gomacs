# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "precise64"
  config.vm.box_url = "http://files.vagrantup.com/precise64.box"

  config.vm.synced_folder `go env GOPATH`.chomp!+"/src", "/home/vagrant/src"

  config.vm.provision "shell", inline: "apt-get update"
  config.vm.provision "shell", inline: "apt-get -y install git emacs"
  config.vm.provision "shell", inline: 'echo "export GOPATH=/home/vagrant" >> .bashrc'
  config.vm.provision "shell", inline: 'echo "export GOROOT=/home/vagrant/go" >> .bashrc'
  config.vm.provision "shell", inline: 'echo "export PATH=$PATH:/home/vagrant/go/bin" >> .bashrc'
  GO_VERSION = "go1.4rc2"
  config.vm.provision "shell", privileged: false, inline: "rm -rf go* && wget -q -O go.linux-amd64.tar.gz https://storage.googleapis.com/golang/"+GO_VERSION+".linux-amd64.tar.gz && tar zxf go.linux-amd64.tar.gz" 
  config.vm.provision "shell", privileged: false, inline: "PS1='$ ' source .bashrc && go get github.com/go-emacs/gomacs && gomacs --install"
end
