# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

overrides = "#{__FILE__}.local"
if File.exist?(overrides)
  eval File.read(overrides)
end

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "bento/debian-9"
  config.vm.box_version = "201812.27.0"

  config.vm.hostname = 'commupace'

  config.vm.network :private_network, ip: "10.1.2.100"
  config.vm.network :forwarded_port, guest: 10000, host: 10000

  config.vm.synced_folder ".", "/vagrant"

  config.vm.provider "virtualbox" do |vb|
    vb.name = "commupace-dev"
    vb.gui = false
    vb.customize ["modifyvm", :id, "--memory", "2048"]
  end

  if Vagrant.has_plugin?("vagrant-proxyconf") && ENV['HTTP_PROXY']
    puts '- Proxy Setting ----------------------------------'
    puts ENV['HTTP_PROXY']
    puts ENV['HTTPS_PROXY']

    config.proxy.http       = ENV['HTTP_PROXY']
    config.proxy.https      = ENV['HTTPS_PROXY']

    config.apt_proxy.http   = ENV['HTTP_PROXY']
    config.apt_proxy.https  = ENV['HTTPS_PROXY']

    config.npm_proxy.http   = ENV['HTTP_PROXY']
    config.npm_proxy.https  = ENV['HTTPS_PROXY']

    config.git_proxy.http   = ENV['HTTP_PROXY']
    config.git_proxy.https  = ENV['HTTPS_PROXY']

    config.proxy.no_proxy = "localhost,127.0.0.1,192.168.0.0/16,172.16.0.0/12,10.0.0.0/8"
    puts '--------------------------------------------------'
  end

  config.vm.provision :shell, :path => 'vagrant-bootstrap.sh'
end
