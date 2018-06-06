# Build
### windows 32
#CGO_ENABLED=0 GOOS=windows GOARCH=386 go build

### windows 64
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build 

### Linux 32
#CGO_ENABLED=0 GOOS=linux GOARCH=386 go build

### Linux 64
cmd=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build;


#######
path=`dirname $0`

cd ${path}/.. && ${cmd}

