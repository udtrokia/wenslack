#!/usr/bin/env sh

path=`dirname $0`

echo ""
read -p "   Please input the server port:  " port
echo""

docker run -d -p ${port}:2333 --restart=always --name wenslack \
       -v `pwd`/${path}/../assets:/go/src/wenslack/assets \
       wenslack ./wenslack


#docker run -p $1:2333  wenslack ./wenslack 


