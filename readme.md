# Zord project creator

This project is a cli to create zord projects easily

## Get cli using installed golang

``` SHELL
go install github.com/levysam/create-zord@latest
```

then export the bin file to your terminal

``` SHELL
export PATH=$PATH:$GOPATH/go/bin/create-zord
```

then run the command

``` SHELL
create-zord create-project
```

## Build cli using raw go build

``` SHELL
go build -o create-zord create-zord.go
```
then you can run the cli:

``` SHELL
./create-zord create-project
```

and follow the interactive terminal

## Build using docker
 
``` SHELL
docker run -v ./:$(pwd) -w $(pwd) golang:1.22-alpine go build -o create-zord create-zord.go
```

then you can run the cli:

``` SHELL
docker run -it -v ./:$(pwd) -w $(pwd) golang:1.22-alpine ./create-zord create-project
```
