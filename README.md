# coreutils

## 简介
The coreutils project is a reimplemented version of golang

## install coreutils
```
env GOPATH=`pwd` go get github.com/guonaihong/coreutils/coreutils
```
If you want to use the cat command
```
./coreutils cat flie
./coreutils cut -d":" -f1 /etc/passwd
./coreutils echo "hello china"
```

## install Compile command separately
```
env GOPATH=`pwd` go run github.com/guonaihong/coreutils/buildall
```
If you want to use the cat command
```
./cat flie
./cut -d":" -f1 /etc/passwd
./echo "hello china"
```

## The completed command is as follows
* basename [detail](./basename/README.md)
* cat [detail](./cat/README.md)
* cut [detail](./cut/README.md)
* dirname [detail](./dirname/README.md)
* echo [detail](./echo/README.md)
* head [detail](./head/README.md)
* paste [detail](./paste/README.md)
* rmdir [detail](./rmdir/README.md)
* tee [detail](./tee/README.md)
* touch [detail](./touch/README.md)
* tail [detail](./tail/README.md)
* tac [detail](./tac/README.md)
* tr [detail](./tr/README.md)
* true [detail](./true/README.md)
* uniq [detail](./uniq/README.md)
* whoami [detail](./whoami/README.md)
* yes [detail](./yes/README.md)
* shuf [detail](./shuf/README.md)
* seq [detail](./seq/README.md)
* sleep [detail](./sleep/README.md)

## progress
progress = 20 / 86 = 23%
