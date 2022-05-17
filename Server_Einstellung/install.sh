sudo apt-get install cmake gcc clang gdb build-essential git-core -y
sudo apt-get install nginx -y
 sudo cp -u ahrensburg.service /etc/systemd/system/ahrensburg.service
sudo  systemctl enable ahrensburg.service
sudo  systemctl start ahrensburg.service

