export GOROOT=/usr/local/go
export GOPATH=/go

export PATH=$GOPATH/bin:$HOME/.cargo/bin:$GOROOT/bin:$PATH
cd /Server



 cd /Server/ahrensburg.blog
go build
cd /Server/dokument_Daux
daux clear-cache
daux generate
