#!/bin/bash


if [ ! -d /etc/opensovs ] 
then
	sudo mkdir /etc/opensovs
fi

sudo sh -c 'echo "
{
	\"DatabaseConf\": {
		\"Host\": \"localhost\",
		\"Port\": 5432,
		\"User\": \"postgres\",
		\"Password\": \"1234\",
		\"DatabaseName\": \"opensovs\"
	},
	\"ApiPort\": 8080
}" > /etc/opensovs/configuration.json'
