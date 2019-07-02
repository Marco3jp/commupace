#!/bin/bash
DBPASSWORD="space-must-have-community"
GO_VERSION="1.12.6"
NODE_VERSION="10.16.0"

sudo echo "vagrant ALL=NOPASSWD: ALL" >> /etc/sudoers
apt update
apt upgrade -y

# Tools and Dependencies
apt install -y g++ git make pkg-config libuv1
# Softwares
apt install -y mariadb-server mariadb-client nginx

if [ ! -e /opt/node ]; then
    # download node v10 and move /opt
    cd /tmp
    curl https://nodejs.org/dist/v$NODE_VERSION/node-v$NODE_VERSION-linux-x64.tar.xz -# -o node-linux-x64.tar.xz
    tar -xvf node-linux-x64.tar.xz
    rm -rf node-linux-x64.tar.xz
    if [ -e /opt/node ]; then
        rm -rf /opt/node
    fi
    mv node-v$NODE_VERSION-linux-x64 /opt/node
    chmod -R a+x /opt/node/bin
    ln -s /opt/node/bin/node /usr/bin/node
fi

if [ ! -e /opt/go ]; then
    cd /tmp
    curl https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz -# -o linux-go.tar.gz
    tar -xvf linux-go.tar.gz
    rm -rf linux-go.tar.gz
    if [ -e /opt/go ]; then
        rm -rf /opt/go/
    fi
    mv go /opt/
fi

if [ ! -e /home/vagrant/.go/src/github.com/Marco3jp ]; then
	mkdir -p /home/vagrant/.go/src/github.com/Marco3jp

	if [ ! -e /home/vagrant/.go/src/github.com/Marco3jp/commupace ]; then
		ln -s /vagrant /home/vagrant/.go/src/github.com/Marco3jp/commupace
	fi
fi

# set environment variable
export PATH=$PATH:/opt/go/bin
export PATH=$PATH:/opt/node/bin
echo "PATH=$PATH" > /etc/environment
export MYSQL_USER=root
echo MYSQL_USER=root >> /etc/environment
export MYSQL_PASSWORD=$DBPASSWORD
echo MYSQL_PASSWORD=$DBPASSWORD >> /etc/environment
export MYSQL_LOCATE=localhost
echo MYSQL_LOCATE=localhost >> /etc/environment
export MYSQL_OPEN_CONN=1000
echo MYSQL_OPEN_CONN=1000 >> /etc/environment
export MYSQL_IDLE_CONN=700
echo MYSQL_IDLE_CONN=700 >> /etc/environment
export MYSQL_CONN_LIFETIME=10000
echo MYSQL_CONN_LIFETIME=10000 >> /etc/environment
source /etc/environment

sudo rm /etc/profile

# link working directory
su -c "echo \"export GOPATH=\\\$HOME/.go/\" >> ~/.bashrc" vagrant
su -c "export PATH=\$PATH:/opt/go/bin && export GOPATH=\$HOME/.go/" vagrant
su -c "cd /vagrant && make deps" vagrant

if [ ! -e /home/vagrant/back ]; then
    su -c "export GOPATH=\$HOME/.go/ &&  ln -sf \${GOPATH}src/github.com/Marco3jp/commupace ~/back" vagrant
fi

if [ ! -e /home/vagrant/front ]; then
    su -c "export GOPATH=\$HOME/.go/ &&  ln -sf \${GOPATH}src/github.com/Marco3jp/commupace/front ~/front" vagrant
fi

# copy configure
#  mariadb
cp /vagrant/vagrant/50-server.cnf /etc/mysql/mariadb.conf.d
mysql -u root -p$DBPASSWORD < /vagrant/vagrant/recreateRootUser.sql

#  nginx
cp /vagrant/vagrant/server.* /home/vagrant/
cp /vagrant/vagrant/web-server.conf /etc/nginx/sites-available/

if [ ! -e /etc/nginx/sites-enabled/web-server.conf ]; then 
    ln -s /etc/nginx/sites-available/web-server.conf /etc/nginx/sites-enabled/web-server.conf
fi

if [ -e /etc/nginx/sites-enabled/default ]; then
    rm /etc/nginx/sites-enabled/default
fi

#  sshd
cp /vagrant/vagrant/sshd_config /etc/ssh/sshd_config

# adjust systemctl service
systemctl stop nginx
systemctl stop mariadb

systemctl enable nginx
systemctl enable mariadb

systemctl start nginx
systemctl start mariadb
systemctl restart sshd

# npm
su -c "cd ~/front && npm install" vagrant

echo "------------------------------"
echo "     Vagrant up success!!"
echo "  Thank you for contribution"
echo "------------------------------"
