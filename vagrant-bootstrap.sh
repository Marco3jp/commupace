#!/bin/bash
DBPASSWORD="space-must-have-community"

apt update
apt upgrade -y

# Tools
apt install -y g++ git make pkg-config
# Use softwares
apt install -y mariadb-server mariadb-client nginx

# download php
if [ ! -e /opt/php/bin/php ]; then
  cd /tmp
  curl --http1.1 https://www.php.net/distributions/php-7.2.17.tar.gz -# -o php7.2.tar.gz
  tar -zxf php7.2.tar.gz
  rm -rf php7.2.tar.gz
  cd php-7.2.17/

# php dependencies
  apt install -y libtool automake bison flex libevent-dev libssl-dev libxml2-dev libbz2-dev libcurl4-openssl-dev libwebp-dev libjpeg-dev libpng-dev libxpm-dev libfreetype6-dev libgmp3-dev libreadline-dev libsnmp-dev libtidy-dev libxslt-dev snmp unzip

# make php and install php
  ./configure --prefix=/opt/php --with-libdir=lib64 --with-pic --with-bz2 --with-freetype-dir --with-png-dir --with-xpm-dir --with-gettext --with-gmp --with-iconv --with-jpeg-dir --with-curl --with-webp-dir --with-png-dir --with-openssl --with-pcre-regex --with-zlib --with-layout=GNU --enable-exif --enable-ftp --enable-sockets --with-kerberos --enable-shmop --enable-calendar --with-libxml-dir --with-mhash --with-readline --with-snmp --with-tidy --with-xsl --with-gnu-ld --enable-mysqlnd --with-mysqli=mysqlnd --with-pdo-mysql=mysqlnd --with-mysql-sock=/var/lib/mysql/mysql.sock --enable-mbstring --with-gd --enable-fpm
  make -j3
  make install
fi

if [ ! -e /usr/bin/node ]; then
  # download node v10 and move /opt
  curl --http1.1 https://nodejs.org/dist/v10.15.3/node-v10.15.3-linux-x64.tar.xz -# -o node-v10.15.3-linux-x64.tar.xz
  tar -Jxvf node-v10.15.3-linux-x64.tar.xz
  rm -rf node-v10.15.3-linux-x64.tar.xz
  if [ -e /opt/node/ ]; then
    rm -rf /opt/node
  fi
  mv node-v10.15.3-linux-x64 /opt/node
  chmod -R a+x /opt/node/bin
fi

# set environment variable
export PATH=$PATH:/opt/php/bin
export PATH=$PATH:/opt/php/sbin
export PATH=$PATH:/usr/local/bin
export PATH=$PATH:/opt/node/bin
echo "PATH=$PATH" > /etc/environment
echo "export PATH=$PATH:/opt/php/bin" > /etc/profile.d/php.sh
echo "export PATH=$PATH:/opt/php/sbin" >> /etc/profile.d/php.sh
export MYSQL_USER=root
echo MYSQL_USER=root >>/etc/environment
export MYSQL_PASSWORD=$DBPASSWORD
echo MYSQL_PASSWORD=$DBPASSWORD >>/etc/environment
export MYSQL_LOCATE=localhost
echo MYSQL_LOCATE=localhost >>/etc/environment
export MYSQL_OPEN_CONN=1000
echo MYSQL_OPEN_CONN=1000 >>/etc/environment
export MYSQL_IDLE_CONN=700
echo MYSQL_IDLE_CONN=700 >>/etc/environment
export MYSQL_CONN_LIFETIME=10000
echo MYSQL_CONN_LIFETIME=10000 >>/etc/environment
source /etc/environment

# download composer installer and install composer
cd /tmp
curl --http1.1 https://getcomposer.org/installer -# -o installer
php installer
mv composer.phar /usr/local/bin/composer

# copy configure
#  mariadb
cp /vagrant/vagrant/50-server.cnf /etc/mysql/mariadb.conf.d
mysql -u root -p$DBPASSWORD < /vagrant/vagrant/recreateRootUser.sql

#  php fpm
cp /vagrant/vagrant/php-fpm.conf /opt/php/etc/
cp /vagrant/vagrant/www.conf /opt/php/etc/php-fpm.d/
cp /vagrant/vagrant/php-fpm.service /lib/systemd/system/php-fpm.service

#  nginx
cp /vagrant/vagrant/web-server.conf /etc/nginx/sites-available/

if [ ! -e /etc/nginx/sites-enabled/web-server.conf ]; then 
  ln -s /etc/nginx/sites-available/web-server.conf /etc/nginx/sites-enabled/web-server.conf
fi

if [ -e /etc/nginx/sites-enabled/default ]; then
  rm /etc/nginx/sites-enabled/default
fi

#  sshd
cp /vagrant/vagrant/sshd_config /etc/ssh/sshd_config

# link working directory
if [ ! -e /home/vagrant/commupace ]; then
  ln -s /vagrant /home/vagrant/commupace
fi
if [ ! -e /home/vagrant/backend ]; then
  ln -s /vagrant/back /home/vagrant/backend
fi

if [ ! -e /home/vagrant/frontend ]; then
  ln -s /vagrant/front /home/vagrant/frontend
fi

# set permission
if [ ! -e /var/run/php-fpm/ ]; then
  chmod -R 775 /opt/
  chown -R www-data /opt/
  mkdir /var/run/php-fpm/
  chown -R www-data /var/run/php-fpm/
  chmod a+x /opt/php/sbin/php-fpm
fi

# npm
su -c "cd ~/frontend && npm install" vagrant

# composer
su -c "cd ~/backend && composer install" vagrant

# adjust systemctl service
systemctl stop nginx
systemctl enable php-fpm
systemctl enable nginx
systemctl start php-fpm
systemctl start nginx
systemctl restart mariadb
systemctl restart sshd

echo "--------------------------"
echo "  vagrant up success!!"
echo "--------------------------"
