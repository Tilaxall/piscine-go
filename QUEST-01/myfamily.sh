#!/bin/bash
#export HERO_ID=70
curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --arg HERO_ID "$HERO_ID" ' .[] | select (.id == ($HERO_ID|tonumber)) | .connections.relatives' | sed 's/"//g'
