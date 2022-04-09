#! /bin/bash

rabbit_deb=rabbitmq-server_2.8.7-1_all.deb

wget http://www.rabbitmq.com/releases/rabbitmq-server/v2.8.7/"$rabbit_deb"
dpkg -i "$rabbit_deb"
sudo apt-get -f install
rm "$rabbit_deb"
#rabbitmq wants this dir in ubuntu, but things should work without it exist
sudo mkdir /etc/rabbitmq/rabbitmq.conf.d
sudo rabbitmq-plugins enable rabbitmq_management
#sudo rabbitmqctl start_app