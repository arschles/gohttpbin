#!/bin/bash
wget https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.5.zip
unzip go_appengine_sdk_linux_amd64-1.9.5.zip
mv go_appengine /usr/local
chmod -R 755 /usr/local/go_appengine
touch /home/vagrant/.bashrc
echo "export PATH=/usr/local/go_appengine:$PATH" >> /home/vagrant/.bashrc
echo "export GOPATH=/vagrant/Godeps/_workspace" >> /home/vagrant/.bashrc
