#!/usr/bin/env bash

####################################
## 更新软件源和软件
####################################
yum update && yum upgrade

####################################
## 启用iptables防火墙
####################################

# 停止firewalld
systemctl stop firewalld
# 禁用firewalld（否则重启系统后会再次启动）
systemctl disable firewalld

# 如果没安装则安装下
yum install iptables-services -y
# 重启iptables
systemctl restart iptables
# 设置开机自启
systemctl enable iptables

####################################
## 安装Golang
####################################

yum install wget -y

latest_version=1.20.1

wget https://dl.google.com/go/go$latest_version.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzf go$latest_version.linux-amd64.tar.gz
rm -fr go$latest_version.linux-amd64.tar.gz

echo "export GOROOT=/usr/local/go" >> ~/.bashrc
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "export GOPATH=~/go" >> ~/.bashrc
source ~/.bashrc

go version

####################################
## 安装Docker
####################################

sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine

sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

sudo systemctl start docker
sudo systemctl enable docker

sudo yum install -y docker-compose

####################################
## 安装Supervisor
####################################

yum install -y epel-release

sudo yum -y install supervisor
sudo systemctl enable supervisord
sudo systemctl start supervisord

####################################
## 安装Nginx
####################################

yum install -y nginx
sudo systemctl enable nginx
sudo systemctl restart nginx

####################################
## 安装其他工具
####################################

yum install -y htop
yum install -y git
yum install -y vim
