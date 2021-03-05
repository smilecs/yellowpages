#!/bin/sh

APPNAME=yellowpages
GOOS=linux
GOARCH=amd64
GO_ENV=production
GO111MODULE=on
SERVER_IP=35.205.155.238

echo 'Building the binary locally'
# build the app for linux
GOOS=linux GOARCH=amd64 go build -o ${APPNAME}

echo 'Sending the binary to the remote server'
# send the binary/files to the server
rsync -auv ./$APPNAME root@$SERVER_IP:/usr/local/bin

echo 'Change the permission of the binary remotely'
# change the permission of the binary
ssh root@$SERVER_IP "sudo chown root: /usr/local/bin/$APPNAME; sudo chmod +x /usr/local/bin/$APPNAME"

echo "stop the api service"
# restart the api service remotely
ssh root@$SERVER_IP "sudo systemctl stop yellowpages.service"

# just wait a while... for the app to stop
sleep 3

# run the migration in the server if migrate is specified
if [ "$1" = "migrate" ]; then
  echo 'Migrate the database in the server'
  ssh root@$SERVER_IP 'psql postgres://postgres:postgres@localhost:5432 -c "drop database yellowpages_production"'
  ssh root@$SERVER_IP 'psql postgres://postgres:postgres@localhost:5432 -c "create database yellowpages_production"'
  ssh root@$SERVER_IP "GO_ENV=production JWT_SECRET=yellowpages-jwt-secret-xtryu /usr/local/bin/$APPNAME migrate"
  ssh root@$SERVER_IP "cd /usr/local/bin/ && $2 GO_ENV=production JWT_SECRET=yellowpages-jwt-secret-xtryu /usr/local/bin/$APPNAME task db:seed"
fi
# $2 should be the MAP_API_KEY=blahblah env variable

echo "start the api service"
# restart the api service remotely
ssh root@$SERVER_IP "sudo systemctl start yellowpages.service"

echo 'Deleting the built binary locally'
rm -rf $APPNAME