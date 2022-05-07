sudo apt-get update -y
sudo apt-get upgrade -y
sudo apt-get update -y
sudo apt-get install php -y
sudo apt-get install gnupg2 -y
sudo apt-get install pandoc -y
sudo apt-get install cmake gcc clang gdb build-essential git-core -y
sudo apt-get install composer -y
sudo apt-get install nginx -y
sudo apt install python3 python3-dev git curl python-is-python3 -y
sudo apt install python3-pip -y
sudo composer global require daux/daux.io
 cd /root
if [ -d /usr/local/go/ ] ; then
 echo "Go ist Vorhanden"
 else
 wget https://go.dev/dl/go1.18.linux-amd64.tar.gz
 sudo tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
 echo "Go wird Installiert"
 fi
export GOROOT=/usr/local/go
export GOPATH=/go
export phpcomposer=/$HOME/.composer/vendor/bin
export PATH=$GOPATH/bin:$phpcomposer:$HOME/.cargo/bin:$GOROOT/bin:$PATH

 cd /Server
 git submodule update --init --recursive


cd /Server/dokument_Daux
daux generate
  cd /Server
 sudo cp -u ahrensburg.service /etc/systemd/system/ahrensburg.service
 sudo cp -u server.service /etc/systemd/system/server.service

sudo  systemctl enable ahrensburg.service
sudo  systemctl start ahrensburg.service