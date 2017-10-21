#!/bin/bash

err() {
    echo "Error on line $1"
}

trap 'err $LINENO' ERR

HOST=192.168.1.20
#HOST=sasha123.ddns.ukrtel.net
SERVICE=goHome

deploy=pi\@"$HOST":/usr/local/bin
deploy_home=pi\@"$HOST":/home/pi/$SERVICE
deploy_ui=pi\@"$HOST":/home/pi/"$SERVICE"/vc


export GOOS=linux
export GOARCH=arm
export GOARM=7



ssh pi\@$HOST "mkdir -p /home/pi/$SERVICE" || exit 1
ssh pi\@$HOST "mkdir -p /home/pi/$SERVICE/ui" || exit 1

scp -r ui/build/default $deploy_ui 

if [ "$1" == "all" ]; then
    go build || exit 1
    ssh pi\@$HOST "sudo service $SERVICE stop" || exit 1
    scp $SERVICE $deploy || exit 1
    ssh pi\@$HOST "sudo service $SERVICE start" || exit 1
fi

echo "Success"
