use mysql;
truncate table user;
flush privileges;
grant all privileges on *.* to root@localhost identified by 'space-must-have-community' with grant option;
grant all privileges on *.* to root@'10.0.0.0/255.0.0.0' identified by 'space-must-have-community' with grant option;
grant all privileges on *.* to root@'172.16.0.0/255.240.0.0' identified by 'space-must-have-community' with grant option;
grant all privileges on *.* to root@'192.168.0.0/255.255.0.0' identified by 'space-must-have-community' with grant option;
flush privileges;
