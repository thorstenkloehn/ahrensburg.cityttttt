export GOROOT=/usr/local/go
export GOPATH=/go
export phpcomposer=/$HOME/.composer/vendor/bin
export PATH=$GOPATH/bin:$phpcomposer:$HOME/.cargo/bin:$GOROOT/bin:$PATH
cd /Server

git pull
git submodule update --recursive --remote

 cd /Server/ahrensburg.blog
go build
cd /Server/dokument_Daux
daux clear-cache
daux generate
