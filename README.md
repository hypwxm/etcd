# etcd


export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/www/go
export GOBIN=$GOPATH/bin
rm -rf /www/go/src/etcd
mkdir -p /www/go/src/etcd
cp -r * /www/go/src/etcd
cd /www/go/src/etcd
go build etcd
./etcd