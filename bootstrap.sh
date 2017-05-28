#/bin/bash
apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

apt-get update
apt-get install -y git docker-ce

usermod -a -G docker ubuntu

wget https://cl.ly/1o3N3x0z0a38/metaops-node
chmod +x metaops-node
